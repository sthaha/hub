package service

import (
	"testing"

	"github.com/tektoncd/hub/api/pkg/testutils"
)

func TestMain(m *testing.M) {
	testutils.Run(m)
}

func TestDummy(t *testing.T) {
	t.Logf("Applied migirations")

}
