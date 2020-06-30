package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ()

func TestCategory_List(t *testing.T) {
	LoadFixture(db, "../../fixtures")
	categorySvc := NewCategory(testConfig)

	all, err := categorySvc.List(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, 3, len(all))
	assert.Equal(t, 2, len(all[0].Tags))
	assert.Equal(t, "abc", all[0].Name) // categories are sorted by name
}
