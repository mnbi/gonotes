// gonotes/internal/textrepo/textrepo_test.go
package textrepo

import (
	"path/filepath"
	"testing"

	"github.com/mnbi/gonotes/internal/conf"
	"github.com/mnbi/gonotes/internal/types"
)

func TestNewRepo(t *testing.T) {
	id := 1
	configFile, _ := filepath.Abs("../../testdata/config.yml")
	_ = conf.Init(configFile)

	var expectedType types.RepoType = types.RT_FILESYSTEM
	expectedName := "gonotes_test"
	expectedRoot, _ := filepath.Abs(filepath.Join("../../testdata", expectedName))
	expectedExt := "md"

	repo := NewRepo()

	if repo.Type != expectedType {
		t.Fatalf("tests[%d] - type wrong, expected=%s, got=%s",
			id, expectedType, repo.Type)
	}

	if repo.Name != expectedName {
		t.Fatalf("tests[%d] - name wrong, expected=%s, got=%s",
			id, expectedName, repo.Name)
	}

	if repo.Root != expectedRoot {
		t.Fatalf("tests[%d] - root wrong, expected=%s, got=%s",
			id, expectedRoot, repo.Root)
	}

	if repo.Ext != expectedExt {
		t.Fatalf("tests[%d] - ext wrong, expected=%s, got=%s",
			id, expectedExt, repo.Ext)
	}
}
