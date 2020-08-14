package ttctl

import (
	"fmt"

	cobra "github.com/spf13/cobra"
)

type workspaceSwitchCommandeer struct {
	cmd                 *cobra.Command
	workspaceCommandeer *workspaceCommandeer
}

func newWorkspaceSwitchCommandeer(workspaceCommandeer *workspaceCommandeer) *workspaceSwitchCommandeer {
	commandeer := &workspaceSwitchCommandeer{
		workspaceCommandeer: workspaceCommandeer,
	}

	cmd := &cobra.Command{
		Use:     "workspace",
		Aliases: []string{"ws"},
		Short:   "Switch to specific workspace.",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("here is a promptui integration,show current worksapce user id and wizard to fetch and switch to other sorkspace.")
			return nil
		},
	}

	commandeer.cmd = cmd

	return commandeer
}
