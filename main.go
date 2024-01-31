package main

import (
	"fmt"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"log"
	"os"
)

func debug(format string, v ...interface{}) {
	format = fmt.Sprintf("[debug] %s\n", format)
	log.Output(2, fmt.Sprintf(format, v...))
}

func main() {
	settings := cli.New()
	cfg := new(action.Configuration)
	//settings.SetNamespace("helm-test")

	un := action.NewUninstall(cfg)
	if err := cfg.Init(settings.RESTClientGetter(), "helm-test", os.Getenv("HELM_DRIVER"), debug); err != nil {
		log.Fatal(err)
	}

	_, err := un.Run("test")
	if err != nil {
		log.Fatal(err)
	}

}
