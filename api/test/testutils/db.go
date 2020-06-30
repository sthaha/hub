package testutils

import (
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/tektoncd/hub/api/pkg/app"
	"github.com/tektoncd/hub/api/pkg/db/model"
)

func LoadFixtures(fixtureDir string) error {
	tc, err := Config()
	if err != nil {
		return err
	}

	fixtures, err := testfixtures.New(
		testfixtures.Database(tc.DB().DB()),
		testfixtures.Dialect(app.DBDialect),
		testfixtures.Directory(fixtureDir))
	if err != nil {
		return err
	}

	return fixtures.Load()
}

func ApplyMigration() {
	tc, _ := Config()
	tc.DB().AutoMigrate(model.Category{}, model.Tag{})
}
