package ttctl

import (
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
		Use:           "ws [command]",
		Aliases:       []string{"workspace"},
		Short:         "Config/Switch/Show local workspace",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.AddCommand(
		newWorkspaceInitCommandeer(commandeer).cmd,
		newWorkspaceShowCommandeer(commandeer).cmd,
		newWorkspaceSwitchCommandeer(commandeer).cmd,
	)

	commandeer.cmd = cmd

	return commandeer
}

type workspaceConfig struct {
	UserId       string `json:"user_id"`
	Organization string `json:"organization"`
	Project      string `json:"project"`
}
