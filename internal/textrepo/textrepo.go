// gonotes/internal/textrepo/textrepo.go
package textrepo

import (
	"path/filepath"

	"github.com/mnbi/gonotes/internal/conf"
	"github.com/mnbi/gonotes/internal/timestamp"
	"github.com/mnbi/gonotes/internal/types"
)

// File type
type FileType string

const (
	MARKDOWN  = "markdown"
	PLAINTEXT = "plain_text"
)

var fileExt map[FileType]string = map[FileType]string{
	MARKDOWN:  "md",
	PLAINTEXT: "txt",
}

const favoriteFileType = MARKDOWN

// Repository
type Repository struct {
	Type types.RepoType
	Name string
	Root string
	Ext  string
}

func NewRepo() *Repository {
	repoType := conf.RepoType()
	repoName := conf.RepoName()
	repoRoot := filepath.Join(conf.RepoBase(), repoName)
	repoExt := fileExt[favoriteFileType]

	return &Repository{
		Type: repoType,
		Name: repoName,
		Root: repoRoot,
		Ext:  repoExt,
	}
}

func (repo *Repository) Entries(stampPattern string) *[]timestamp.Timestamp {
	entries := make([]timestamp.Timestamp, 0, 0)
	return &entries
}
