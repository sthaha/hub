package testutils

import (
	"fmt"
	"os"
	"testing"
)

func Run(m *testing.M) {
	tc, err := Config()
	if err != nil {
		fmt.Fprintf(os.Stderr, "FATAL: failed to initialise: %s", err)
		os.Exit(1)
	}
	defer tc.Cleanup()

	ApplyMigration()

	os.Exit(m.Run())
}
