package app

import (
	"github.com/Slahser/coup-de-grace/pkg/ttctl"
	"net/http"
	"github.com/inconshreveable/go-update"
)

func Run() error {
	return ttctl.NewRootCommandeer().Execute()
}

func doUpdate(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := update.Apply(resp.Body, update.Options{}); err != nil {
		// error handling
	}
	return err
}