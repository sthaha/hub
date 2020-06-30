package testutils

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/tektoncd/hub/api/pkg/app"
)

var once sync.Once
var tc app.Config
var err error

var fixturePath string

func Config() (app.Config, error) {
	once.Do(func() {
		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			panic("No caller information")
		}

		fmt.Printf("Filename : %q, Dir : %q\n", filename, path.Dir(filename))

		envPath := filepath.Join(path.Dir(filename), "..", "config", "env.test")
		fixturePath = filepath.Join(path.Dir(filename), "..", "fixtures")
		tc, err = app.FromEnvFile(envPath)
	})

	return tc, err
}

func FixturePath() string {
	return fixturePath
}
