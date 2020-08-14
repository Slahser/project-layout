package ttctl

import (
	cobra "github.com/spf13/cobra"
)

type funcCommandeer struct {
	cmd            *cobra.Command
	rootCommandeer *RootCommandeer
}

func newFuncCommandeer(rootCommandeer *RootCommandeer) *funcCommandeer {
	commandeer := &funcCommandeer{
		rootCommandeer: rootCommandeer,
	}

	cmd := &cobra.Command{
		Use:           "func [command]",
		Aliases:       []string{"function"},
		Short:         "Operations for Tt Function.",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.AddCommand(
		newFuncCreateCommandeer(commandeer).cmd,
		newFuncBuildImageCommandeer(commandeer).cmd,
		newFuncDescribeCommandeer(commandeer).cmd,
		newFuncDeleteCommandeer(commandeer).cmd,
		newFuncUpdateCommandeer(commandeer).cmd,
	)

	commandeer.cmd = cmd

	return commandeer
}
