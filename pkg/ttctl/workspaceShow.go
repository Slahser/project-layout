package ttctl

import (
	"encoding/json"
	"fmt"

	cobra "github.com/spf13/cobra"
)

type workspaceShowCommandeer struct {
	cmd                 *cobra.Command
	workspaceCommandeer *workspaceCommandeer
}

func newWorkspaceShowCommandeer(workspaceCommandeer *workspaceCommandeer) *workspaceShowCommandeer {
	commandeer := &workspaceShowCommandeer{
		workspaceCommandeer: workspaceCommandeer,
	}

	cmd := &cobra.Command{
		Use:     "workspace",
		Aliases: []string{"ws"},
		Short:   "Switch to specific workspace.",
		RunE: func(cmd *cobra.Command, args []string) error {
			show()
			return nil
		},
	}

	commandeer.cmd = cmd

	return commandeer
}

func show() {
	mockedWorkspaceConfig := &workspaceConfig{"0x3000001", "mytt-org", "mytt-project"}

	serializedConfig, _ := json.MarshalIndent(mockedWorkspaceConfig, "", "    ")
	fmt.Printf(string(serializedConfig))
}
