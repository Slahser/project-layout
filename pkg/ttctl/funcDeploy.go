package ttctl

import (
	"fmt"

	"github.com/spf13/cobra"
)

type funcDeployCommandeer struct {
	cmd            *cobra.Command
	funcCommandeer *funcCommandeer
}

func newFuncDeployCommandeer(funcCommandeer *funcCommandeer) *funcDeployCommandeer {
	commandeer := &funcDeployCommandeer{
		funcCommandeer: funcCommandeer,
	}

	cmd := &cobra.Command{
		Use:     "func",
		Aliases: []string{"function"},
		Short:   "Deploy a Function ",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Client version")
			return nil
		},
	}

	commandeer.cmd = cmd

	return commandeer
}
