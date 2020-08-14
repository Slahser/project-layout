package main

import "github.com/Slahser/coup-de-grace/pkg/ttctl"

func main() {
	_ = ttctl.NewRootCommandeer().Execute()
}
