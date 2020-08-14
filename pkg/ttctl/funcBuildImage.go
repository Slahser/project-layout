package ttctl

import (
	"fmt"

	"github.com/spf13/cobra"
)

type funcBuildImageCommandeer struct {
	cmd            *cobra.Command
	funcCommandeer *funcCommandeer
}

func newFuncBuildImageCommandeer(funcCommandeer *funcCommandeer) *funcBuildImageCommandeer {
	commandeer := &funcBuildImageCommandeer{
		funcCommandeer: funcCommandeer,
	}

	cmd := &cobra.Command{
		Use:     "build-image [-flag]",
		Aliases: []string{"build"},
		Short:   "Build a function image.",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Client version")
			return nil
		},
	}

	commandeer.cmd = cmd

	return commandeer
}
