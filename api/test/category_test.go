package hub

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/ikawaha/goahttpcheck"
	"github.com/stretchr/testify/assert"
	category "github.com/tektoncd/hub/api/gen/category"
	categoryServer "github.com/tektoncd/hub/api/gen/http/category/server"
	"github.com/tektoncd/hub/api/pkg/service"
	"github.com/tektoncd/hub/api/test/testutils"
)

func TestCategories_List(t *testing.T) {
	tc, _ := testutils.Config()

	if err := LoadFixtures(tc, "./fixtures"); err != nil {
		assert.FailNow(t, "Failed to load fixtures", err)
	}

	checker := goahttpcheck.New()
	checker.Mount(
		categoryServer.NewListHandler,
		categoryServer.MountListHandler,
		category.NewListEndpoint(service.NewCategory(tc)))

	checker.Test(t, http.MethodGet, "/categories").
		Check().HasStatus(http.StatusOK).Cb(func(r *http.Response) {

		b, err := ioutil.ReadAll(r.Body)
		assert.NoError(t, err)
		defer r.Body.Close()

		var jsonMap []map[string]interface{}
		err = json.Unmarshal([]byte(b), &jsonMap)
		assert.NoError(t, err)

		assert.Equal(t, 3, len(jsonMap))
		assert.Equal(t, "abc", jsonMap[0]["name"])
	})
}
