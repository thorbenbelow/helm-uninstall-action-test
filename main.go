package main

import (
	"fmt"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"log"
	"os"
)

var settings = cli.New()

func debug(format string, v ...interface{}) {
	if settings.Debug {
		format = fmt.Sprintf("[debug] %s\n", format)
		log.Output(2, fmt.Sprintf(format, v...))
	}
}

func main() {
	cfg := new(action.Configuration)

	// Create a new uninstall action
	un := action.NewUninstall(cfg)

	// Initialize the config with the release namespace
	if err := cfg.Init(settings.RESTClientGetter(), "helm-test", os.Getenv("HELM_DRIVER"), debug); err != nil {
		log.Fatal(err)
	}

	// Uninstall the release
	_, err := un.Run("test")
	if err != nil {
		log.Fatal(err)
	}
}
