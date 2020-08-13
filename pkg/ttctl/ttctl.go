package ttctl

import (
	cobra "github.com/spf13/cobra"
	doc "github.com/spf13/cobra/doc"
	zap "go.uber.org/zap"
)

type RootCommandeer struct {
	logger  *zap.SugaredLogger
	cmd     *cobra.Command
	verbose bool
}

func NewRootCommandeer() *RootCommandeer {
	commandeer := &RootCommandeer{}

	cmd := &cobra.Command{
		Use:           "ttctl [command]",
		Short:         "Tt terminal ui",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.PersistentFlags().BoolVarP(&commandeer.verbose, "verbose", "v", false, "Verbose mode, show more details")

	cmd.AddCommand(
		//newBuildCommandeer(commandeer).cmd,
		//newDeployCommandeer(commandeer).cmd,
		//newInvokeCommandeer(commandeer).cmd,
		//newGetCommandeer(commandeer).cmd,
		//newDeleteCommandeer(commandeer).cmd,
		//newUpdateCommandeer(commandeer).cmd,
		//newCreateCommandeer(commandeer).cmd,
		//newExportCommandeer(commandeer).cmd,
		newVersionCommandeer(commandeer).cmd,
		newWorkspaceCommandeer(commandeer).cmd,
		newFuncCommandeer(commandeer).cmd,
	)

	commandeer.cmd = cmd

	return commandeer
}

// Execute uses os.Args to execute the command
func (rc *RootCommandeer) Execute() error {
	return rc.cmd.Execute()
}

// GetCmd returns the underlying cobra command
func (rc *RootCommandeer) GetCmd() *cobra.Command {
	return rc.cmd
}

// CreateMarkdown generates MD files in the target path
func (rc *RootCommandeer) CreateMarkdown(path string) error {
	return doc.GenMarkdownTree(rc.cmd, path)
}

func (rc *RootCommandeer) initialize() error {
	var err error

	rc.logger, err = rc.createLogger()
	if err != nil {
		return err
	}

	return nil
}

func (rc *RootCommandeer) createLogger() (*zap.SugaredLogger, error) {

	var loggingConfig zap.Config

	if rc.verbose {
		loggingConfig = zap.NewDevelopmentConfig()
	} else {
		loggingConfig = zap.NewProductionConfig()
	}

	logger, err := loggingConfig.Build()
	if logger != nil {
		sugar := logger.Sugar()
		return sugar, nil
	} else {
		return nil, err
	}
}
