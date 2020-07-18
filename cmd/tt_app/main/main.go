package main

import (
	"bytes"
	"fmt"
	"github.com/Slahser/coup-de-grace/internal/app/tt_practice"
	"github.com/dimiro1/banner"
	_ "github.com/dimiro1/banner/autoload"
	"github.com/felixge/fgprof"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

func main() {

	http.DefaultServeMux.Handle("/debug/fgprof", fgprof.Handler())
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	nyanFilePath := "/Users/Slahser/Documents/repository/go/pkg/mod/github.com/dimiro1/banner@v1.0.0/example/nyancat.txt"
	nyanBuf, err := ioutil.ReadFile(nyanFilePath)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
	}

	isEnabled := true
	isColorEnabled := true
	banner.Init(os.Stdout, isEnabled, isColorEnabled, bytes.NewBuffer(nyanBuf))

	//=====
	start := time.Now()

	//practice
	tt_practice.Tt1()
	tt_practice.Tt2()
	tt_practice.Tt3()
	tt_practice.Tt4()
	tt_practice.Tt5()
	tt_practice.Tt6()
	//=====
	end := time.Now()
	cost := end.Sub(start)
	fmt.Printf("total cost is %s\n", cost)
}
