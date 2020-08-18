package ttctl

import (
	"fmt"
	"io/ioutil"
	"os"

	homedir "github.com/mitchellh/go-homedir"
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
		Use:   "show",
		Short: "Show current workspace.",
		RunE: func(cmd *cobra.Command, args []string) error {
			show()
			return nil
		},
	}

	commandeer.cmd = cmd

	return commandeer
}

func show() {

	basicPath, _ := homedir.Dir()
	inputFile, inputError := os.Open(basicPath + "/.ttctl/workspace.json")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile")
		return
	}
	defer inputFile.Close()

	inputBytes, _ := ioutil.ReadAll(inputFile)

	var wsConfig workspaceConfig

	_ = json.Unmarshal(inputBytes, &wsConfig)

	serializedConfig, _ := json.MarshalIndent(wsConfig, "", "    ")
	fmt.Print(string(serializedConfig))
}
