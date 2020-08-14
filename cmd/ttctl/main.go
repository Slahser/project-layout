package main

import (
	"fmt"
	"io/ioutil"
	"os"

	jsoniter "github.com/json-iterator/go"

	"github.com/mitchellh/go-homedir"

	"github.com/Slahser/coup-de-grace/cmd/ttctl/app"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func main() {

	show()

	if err := app.Run(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
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
	fmt.Printf(string(serializedConfig))
}

type workspaceConfig struct {
	UserId       string `json:"user_id"`
	Organization string `json:"organization"`
	Project      string `json:"project"`
}
