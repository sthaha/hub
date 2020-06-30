package testutils

import (
	"sync"

	"github.com/tektoncd/hub/api/pkg/app"
)

var once sync.Once
var tc app.Config
var err error

func Config() (app.Config, error) {
	once.Do(func() {
		tc, err = app.FromEnvFile("./config/env.test")
	})
	return tc, err
}
