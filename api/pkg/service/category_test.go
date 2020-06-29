package hub

import (
	"context"
	"github.com/stretchr/testify/assert"
	category "github.com/tektoncd/hub/api/gen/category"
	"testing"
)

var (
	categorySvc category.Service
)

func Test_All(t *testing.T) {
	LoadFixture(db, "../../fixtures")
	categorySvc = NewCategory(testConfig)
	all, err := categorySvc.All(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, 3, len(all))
	assert.Equal(t, 2, len(all[0].Tags))
	assert.Equal(t, "abc", all[0].Name) // categories are sorted by name
}
