package hub

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/tektoncd/hub/api/pkg/app"
	"github.com/tektoncd/hub/api/pkg/db/model"
	"github.com/tektoncd/hub/api/test/testutils"
)

func LoadFixtures(conf app.Config, fixtureDir string) error {
	fixtures, err := testfixtures.New(
		testfixtures.Database(conf.DB().DB()),
		testfixtures.Dialect(app.DBDialect),
		testfixtures.Directory(fixtureDir),
	)

	if err != nil {
		return err
	}

	return fixtures.Load()
}

func TestMain(m *testing.M) {
	tc, err := testutils.Config()
	if err != nil {
		fmt.Fprintf(os.Stderr, "FATAL: failed to initialise: %s", err)
		os.Exit(1)
	}

	tc.DB().AutoMigrate(model.Category{}, model.Tag{})

	defer os.Exit(m.Run())
	defer tc.Cleanup()
}
