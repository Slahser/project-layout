package ttctl

import (
	"fmt"

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
		Use:     "func",
		Aliases: []string{"function"},
		Short:   "Operations for Tt Function.",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Client version")
			return nil
		},
	}

	commandeer.cmd = cmd

	return commandeer
}
