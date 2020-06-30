package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tektoncd/hub/api/test/testutils"
)

func TestCategory_List(t *testing.T) {
	if err := testutils.LoadFixtures(testutils.FixturePath()); err != nil {
		assert.FailNow(t, "Failed to load fixtures", err)
	}

	tc, _ := testutils.Config()
	categorySvc := NewCategory(tc)
	all, err := categorySvc.List(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, 3, len(all))
	assert.Equal(t, 2, len(all[0].Tags))
	assert.Equal(t, "abc", all[0].Name) // categories are sorted by name
}
