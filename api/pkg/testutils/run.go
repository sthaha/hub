package testutils

import (
	"fmt"
	"os"
	"testing"
)

func Run(m *testing.M) {
	tc := Config()
	if !tc.IsValid() {
		fmt.Fprintf(os.Stderr, "Failed to create test configuration object: %s", tc.Error())
		os.Exit(1)
	}
	defer tc.Cleanup()

	if err := applyMigration(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to apply migration: %s", tc.Error())
		os.Exit(1)
	}

	os.Exit(m.Run())
}
