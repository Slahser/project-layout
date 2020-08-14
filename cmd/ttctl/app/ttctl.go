package app

import "github.com/Slahser/coup-de-grace/pkg/ttctl"

func Run() error {
	return ttctl.NewRootCommandeer().Execute()
}
