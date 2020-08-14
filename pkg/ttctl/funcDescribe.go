package ttctl

import (
	"fmt"

	"github.com/spf13/cobra"
)

type funcDescribeCommandeer struct {
	cmd            *cobra.Command
	funcCommandeer *funcCommandeer
}

func newFuncDescribeCommandeer(funcCommandeer *funcCommandeer) *funcDescribeCommandeer {
	commandeer := &funcDescribeCommandeer{
		funcCommandeer: funcCommandeer,
	}

	cmd := &cobra.Command{
		Use:     "desc [-flag]",
		Aliases: []string{"describe"},
		Short:   "Describe specific Function.",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Client version")
			return nil
		},
	}

	commandeer.cmd = cmd

	return commandeer
}
