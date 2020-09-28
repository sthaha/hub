package parser

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/tektoncd/hub/api/pkg/git"
	"go.uber.org/zap"
	"gotest.tools/assert"
	"gotest.tools/assert/cmp"
)

type fakeRepo struct {
	path         string
	head         string
	modifiedTime map[string]time.Time
}

var _ git.Repo = (*fakeRepo)(nil)

func (r fakeRepo) Path() string {
	return r.path
}

func (r fakeRepo) Head() string {
	return r.head
}

func (r fakeRepo) ModifiedTime(path string) (time.Time, error) {

	rp, _ := r.RelPath(path)
	return r.modifiedTime[rp], nil
}

func (r fakeRepo) RelPath(f string) (string, error) {
	return filepath.Rel(r.path, f)
}

func TestParse_NonExistentRepo(t *testing.T) {
	repo := fakeRepo{
		path: "./testdata/catalogs/non-existent",
	}

	p := ForCatalog(zap.NewNop().Sugar(), repo, "")
	_, result := p.Parse()
	assert.Equal(t, 1, len(result.Errors))
	assert.Equal(t, "no resources found in repo", result.Error())

}

func TestParse_ValidRepo(t *testing.T) {
	now := time.Now()
	repo := fakeRepo{
		path: "./testdata/catalogs/valid",
		modifiedTime: map[string]time.Time{
			"task/maven/0.1/maven.yaml": now,
		},
	}

	p := ForCatalog(zap.NewNop().Sugar(), repo, "")
	res, result := p.Parse()

	assert.Equal(t, 0, len(result.Errors))
	assert.Equal(t, "", result.Error())

	assert.Equal(t, 2, len(res))

	gitCLI := res[0]
	assert.Equal(t, "git-cli", gitCLI.Name)
	assert.Equal(t, 1, len(gitCLI.Versions))

	maven := res[1]
	assert.Equal(t, "maven", maven.Name)
	assert.Equal(t, 2, len(maven.Versions))
	assert.Equal(t, now, maven.Versions[0].ModifiedAt)
}

func TestParse_InvalidTask(t *testing.T) {
	// invalid task is ignored but result must have the issue it found
	repo := fakeRepo{
		path: "./testdata/catalogs/invalid-task",
	}

	p := ForCatalog(zap.NewNop().Sugar(), repo, "")
	res, result := p.Parse()

	assert.Equal(t, 0, len(res))

	assert.Equal(t, 1, len(result.Errors))
	assert.Equal(t, "no resources found in repo", result.Error())

	assert.Equal(t, 1, len(result.Issues))
	issue := result.Issues[0]
	assert.Assert(t, cmp.Contains(issue.Message, "git-cli is missing mandatory version label"))
	assert.Equal(t, Critical, issue.Type)
}
