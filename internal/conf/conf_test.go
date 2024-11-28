// gonotes/internal/conf/conf_test.go
package conf

import (
	"path/filepath"
	"testing"

	"github.com/mnbi/gonotes/internal/types"
)

func TestConf(t *testing.T) {
	id := 1

	if RepoType() != types.RT_FILESYSTEM {
		t.Fatalf("tests[%d] - RepoType wrong, expected=%s, got=%s",
			id, types.RT_FILESYSTEM, RepoType())
	}
}

func TestInit(t *testing.T) {
	id := 100
	configFile, _ := filepath.Abs("../../testdata/config.yml")
	_ = Init(configFile)

	expectedName := "gonotes_test"
	expectedBase, _ := filepath.Abs("../../testdata")

	if RepoType() != types.RT_FILESYSTEM {
		t.Fatalf("tests[%d] - RepoType wrong, expected=%s, got=%s",
			id, types.RT_FILESYSTEM, RepoType())
	}

	if RepoName() != expectedName {
		t.Fatalf("tests[%d] - RepoName wrong, expected=%s, got=%s",
			id, expectedName, RepoName())
	}

	if RepoBase() != expectedBase {
		t.Fatalf("tests[%d] - RepoBase wrong, expected=%s, got=%s",
			id, expectedBase, RepoBase())
	}
}
