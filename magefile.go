// +build mage
//https://magefile.org/
//https://github.com/magefile/awesome-mage
package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// allow user to override go executable by running as GOEXE=xxx make ... on unix-like systems
var goexe = "go"
var docker = sh.RunCmd("docker")

var (
	packageName  = "github.com/Slahser/coup-de-grace"
	pkgPrefixLen = len("github.com/Slahser/coup-de-grace")
	noGitLdflags = "-X $PACKAGE/common/slahser.buildDate=$BUILD_DATE"
	ldflags      = "-X $PACKAGE/common/slahser.commitHash=$COMMIT_HASH "+ noGitLdflags
	pkgs         []string
	pkgsInit     sync.Once
)

func init() {
	if exe := os.Getenv("GOEXE"); exe != "" {
		goexe = exe
	}

	// We want to use Go 1.11 modules even if the source lives inside GOPATH.
	// The default is "auto".
	_ = os.Setenv("GO111MODULE", "on")
}

// Install binary
func Install() error {
	return runWith(flagEnv(), goexe, "install", "-ldflags", ldflags, buildFlags(), "-tags", buildTags(), packageName)
}

// Uninstall binary
func Uninstall() error {
	return sh.Run(goexe, "clean", "-i", packageName)
}

// GenDocsHelper Gen docs helper
func GenDocsHelper() error {
	return runCmd(flagEnv(), goexe, "run", "-tags", buildTags(), "main.go", "gen", "docshelper")
}

// Check  tests and linters
func Check() {
	if strings.Contains(runtime.Version(), "1.8") {
		// Go 1.8 doesn't play along with go test ./... and /vendor.
		// We could fix that, but that would take time.
		fmt.Printf("Skip Check on %s\n", runtime.Version())
		return
	}

	if runtime.GOARCH == "amd64" && runtime.GOOS != "darwin" {
		mg.Deps(Test386)
	} else {
		fmt.Printf("Skip Test386 on %s and/or %s\n", runtime.GOARCH, runtime.GOOS)
	}

	mg.Deps(Fmt, Vet)

	// don't run two tests in parallel, they saturate the CPUs anyway, and running two
	// causes memory issues in CI.
	mg.Deps(TestRace)
}

// Build hugo Docker container
func Docker() error {
	if err := docker("build", "-t", "hugo", "."); err != nil {
		return err
	}
	// yes ignore errors here
	_ = docker("rm", "-f", "hugo-build")
	if err := docker("run", "--name", "hugo-build", "hugo ls /go/bin"); err != nil {
		return err
	}
	if err := docker("cp", "hugo-build:/go/bin/hugo", "."); err != nil {
		return err
	}
	return docker("rm", "hugo-build")
}

// Run tests
func Test() error {
	env := map[string]string{"GOFLAGS": testGoFlags()}
	return runCmd(env, goexe, "test", "./...", buildFlags(), "-tags", buildTags())
}

// Run tests with race detector
func TestRace() error {
	env := map[string]string{"GOFLAGS": testGoFlags()}
	return runCmd(env, goexe, "test", "-race", "./...", buildFlags(), "-tags", buildTags())
}

// Run tests in 32-bit mode
// Note that we don't run with the extended tag. Currently not supported in 32 bit.
func Test386() error {
	env := map[string]string{"GOARCH": "386", "GOFLAGS": testGoFlags()}
	return runCmd(env, goexe, "test", "./...")
}

// Run gofmt linter
func Fmt() error {
	if !isGoLatest() {
		return nil
	}
	pkgs, err := currentPackages()
	if err != nil {
		return err
	}
	failed := false
	first := true
	for _, pkg := range pkgs {
		files, err := filepath.Glob(filepath.Join(pkg, "*.go"))
		if err != nil {
			return nil
		}
		for _, f := range files {
			// gofmt doesn't exit with non-zero when it finds unformatted code
			// so we have to explicitly look for output, and if we find any, we
			// should fail this target.
			s, err := sh.Output("gofmt", "-l", f)
			if err != nil {
				fmt.Printf("ERROR: running gofmt on %q: %v\n", f, err)
				failed = true
			}
			if s != "" {
				if first {
					fmt.Println("The following files are not gofmt'ed:")
					first = false
				}
				failed = true
				fmt.Println(s)
			}
		}
	}
	if failed {
		return errors.New("improperly formatted go files")
	}
	return nil
}

// Run golint linter
func Lint() error {
	pkgs, err := currentPackages()
	if err != nil {
		return err
	}
	failed := false
	for _, pkg := range pkgs {
		// We don't actually want to fail this target if we find golint errors,
		// so we don't pass -set_exit_status, but we still print out any failures.
		if _, err := sh.Exec(nil, os.Stderr, nil, "golint", pkg); err != nil {
			fmt.Printf("ERROR: running go lint on %q: %v\n", pkg, err)
			failed = true
		}
	}
	if failed {
		return errors.New("errors running golint")
	}
	return nil
}

//  Run go vet linter
func Vet() error {
	if err := sh.Run(goexe, "vet", "./..."); err != nil {
		return fmt.Errorf("error running go vet: %v", err)
	}
	return nil
}

// Generates a new release.  Expects the TAG environment variable to be set,
// which will create a new tag with that name.
func Release() (err error) {
	releaseTag := regexp.MustCompile(`^v1\.[0-9]+\.[0-9]+$`)
	//https://goreleaser.com/intro/
	tag := os.Getenv("TAG")
	if !releaseTag.MatchString(tag) {
		return errors.New("TAG environment variable must be in semver v1.x.x format, but was " + tag)
	}

	if err := sh.RunV("git", "tag", "-a", tag, "-m", tag); err != nil {
		return err
	}
	if err := sh.RunV("git", "push", "origin", tag); err != nil {
		return err
	}
	defer func() {
		if err != nil {
			sh.RunV("git", "tag", "--delete", "$TAG")
			sh.RunV("git", "push", "--delete", "origin", "$TAG")
		}
	}()
	return sh.RunV("goreleaser")
}

// Generate test coverage report
func TestCoverHTML() error {
	const (
		coverAll = "coverage-all.out"
		cover    = "coverage.out"
	)
	f, err := os.Create(coverAll)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write([]byte("mode: count")); err != nil {
		return err
	}
	pkgs, err := currentPackages()
	if err != nil {
		return err
	}
	for _, pkg := range pkgs {
		if err := sh.Run(goexe, "test", "-coverprofile="+cover, "-covermode=count", pkg); err != nil {
			return err
		}
		b, err := ioutil.ReadFile(cover)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return err
		}
		idx := bytes.Index(b, []byte{'\n'})
		b = b[idx+1:]
		if _, err := f.Write(b); err != nil {
			return err
		}
	}
	if err := f.Close(); err != nil {
		return err
	}
	return sh.Run(goexe, "tool", "cover", "-html="+coverAll)
}

func flagEnv() map[string]string {
	hash, _ := sh.Output("git", "rev-parse", "--short", "HEAD")
	return map[string]string{
		"PACKAGE":     packageName,
		"COMMIT_HASH": hash,
		"BUILD_DATE":  time.Now().Format("2006-01-02T15:04:05Z0700"),
	}
}

func argsToStrings(v ...interface{}) []string {
	var args []string
	for _, arg := range v {
		switch v := arg.(type) {
		case string:
			if v != "" {
				args = append(args, v)
			}
		case []string:
			if v != nil {
				args = append(args, v...)
			}
		default:
			panic("invalid type")
		}
	}

	return args
}

func isGoLatest() bool {
	return strings.Contains(runtime.Version(), "1.14")
}

func runCmd(env map[string]string, cmd string, args ...interface{}) error {
	if mg.Verbose() {
		return runWith(env, cmd, args...)
	}
	output, err := sh.OutputWith(env, cmd, argsToStrings(args...)...)
	if err != nil {
		fmt.Fprint(os.Stderr, output)
	}

	return err
}

func runWith(env map[string]string, cmd string, inArgs ...interface{}) error {
	s := argsToStrings(inArgs...)
	return sh.RunWith(env, cmd, s...)

}

func currentPackages() ([]string, error) {
	var err error
	pkgsInit.Do(func() {
		var s string
		s, err = sh.Output(goexe, "list", "./...")
		if err != nil {
			return
		}
		pkgs = strings.Split(s, "\n")
		for i := range pkgs {
			pkgs[i] = "." + pkgs[i][pkgPrefixLen:]
		}
	})
	return pkgs, err
}

func testGoFlags() string {
	if isCI() {
		return ""
	}

	return "-test.short"
}

func isCI() bool {
	return os.Getenv("CI") != ""
}

func buildFlags() []string {
	if runtime.GOOS == "windows" {
		return []string{"-buildmode", "exe"}
	}
	return nil
}

func buildTags() string {
	// To build the extended Hugo SCSS/SASS enabled version, build with
	// SLAHSER_BUILD_TAGS=extended mage install etc.
	if envtags := os.Getenv("SLAHSER_BUILD_TAGS"); envtags != "" {
		return envtags
	}
	return "none"
}

func xplatPath(pathParts ...string) string {
	return filepath.Join(pathParts...)
}

func gitCommit(shortVersion bool) (string, error) {
	args := []string{
		"rev-parse",
	}
	if shortVersion {
		args = append(args, "--short")
	}
	args = append(args, "HEAD")
	val, valErr := sh.Output("git", args...)
	return strings.TrimSpace(val), valErr
}
