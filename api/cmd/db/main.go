package main

import (
	"fmt"
	"os"

	app "github.com/tektoncd/hub/api/pkg/app"
	"github.com/tektoncd/hub/api/pkg/db/model"
)

func main() {
	api, err := app.FromEnv()
	if err != nil {
		fmt.Fprintf(os.Stderr, "FATAL: failed to initialise: %s", err)
		os.Exit(1)
	}
	defer api.Cleanup()

	logger := api.Logger()
	if err = model.Migrate(api); err != nil {
		logger.Infof("DB initialisation failed", err)
	}
	logger.Info("DB initialisation successful")
}
