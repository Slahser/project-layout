package main

import (
	"os"

	jsoniter "github.com/json-iterator/go"

	"github.com/Slahser/coup-de-grace/cmd/ttctl/app"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func main() {

	if err := app.Run(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
