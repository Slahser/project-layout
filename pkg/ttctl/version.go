package ttctl

import (
	"fmt"

	cobra "github.com/spf13/cobra"
	version "github.com/v3io/version-go"
)

type versionCommandeer struct {
	cmd            *cobra.Command
	rootCommandeer *RootCommandeer
}

func newVersionCommandeer(rootCommandeer *RootCommandeer) *versionCommandeer {
	commandeer := &versionCommandeer{
		rootCommandeer: rootCommandeer,
	}

	cmd := &cobra.Command{
		Use:     "version",
		Aliases: []string{"ver"},
		Short:   "Display the version of the ttctl CLI",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Client version:\n%#v", version.Get().String())
			return nil
		},
	}

	commandeer.cmd = cmd

	return commandeer
}
