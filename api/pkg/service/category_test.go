package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tektoncd/hub/api/pkg/testutils"
)

func TestCategory_List(t *testing.T) {
	tc := testutils.Config()
	testutils.LoadFixtures(t, tc.FixturePath())

	categorySvc := NewCategory(tc)
	all, err := categorySvc.List(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, 3, len(all))
	assert.Equal(t, 2, len(all[0].Tags))
	assert.Equal(t, "abc", all[0].Name) // categories are sorted by name
}
