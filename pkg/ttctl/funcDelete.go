package ttctl

import (
	"fmt"

	"github.com/spf13/cobra"
)

type funcDeleteCommandeer struct {
	cmd            *cobra.Command
	funcCommandeer *funcCommandeer
}

func newFuncDeleteCommandeer(funcCommandeer *funcCommandeer) *funcDeleteCommandeer {
	commandeer := &funcDeleteCommandeer{
		funcCommandeer: funcCommandeer,
	}

	cmd := &cobra.Command{
		Use:     "delete [-flag]",
		Aliases: []string{"remove", "uninstall"},
		Short:   "Delete a Function.",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Client version")
			return nil
		},
	}

	commandeer.cmd = cmd

	return commandeer
}
