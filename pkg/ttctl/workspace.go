package ttctl

import (
	"fmt"

	cobra "github.com/spf13/cobra"
)

type workspaceCommandeer struct {
	cmd            *cobra.Command
	rootCommandeer *RootCommandeer
}

func newWorkspaceCommandeer(rootCommandeer *RootCommandeer) *workspaceCommandeer {
	commandeer := &workspaceCommandeer{
		rootCommandeer: rootCommandeer,
	}

	cmd := &cobra.Command{
		Use:     "workspace",
		Aliases: []string{"ws"},
		Short:   "Switch to specific workspace.",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Client version")
			return nil
		},
	}

	commandeer.cmd = cmd

	return commandeer
}
