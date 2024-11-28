// gonotes/internal/conf/conf.go

// Provides interfaces to refer the configuration settings.
package conf

import (
	"bufio"
	"log"
	"maps"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/mnbi/gonotes"
	"github.com/mnbi/gonotes/internal/env"
	"github.com/mnbi/gonotes/internal/types"
)

const (
	repoType = ":repository_type"
	repoName = ":repository_name"
	repoBase = ":repository_base"
)

var confValueMap map[string]types.RepoType = map[string]types.RepoType{
	":file_system": types.RT_FILESYSTEM,
}

var confMap map[string]string

func init() {
	confMap = defaultConfMap()
}

func defaultConfMap() map[string]string {
	cm := make(map[string]string)
	cm[repoType] = types.RT_FILESYSTEM
	cm[repoName] = "notes"
	cm[repoBase] = "~"
	return cm
}

// Initialize the conf package. This function must be called before
// any other functions are called. When the argument is an empty
// string, the default configuration file will be loaded.
//
// Even if the specified file does not exist, the default values are
// prepared.
func Init(conffile string) error {
	var err error

	if conffile == "" {
		conffile = defaultConfFile()
	}
	if _, err := os.Stat(conffile); err != nil {
		log.Printf("not found the configuration file: %s\n", conffile)
		return err
	} else {
		yamlMap := loadConfYaml(conffile)
		maps.Copy(confMap, yamlMap)
	}

	// expand the base directory to the absolute path
	base := confMap[repoBase]
	if strings.Contains(base, "~") {
		base = strings.Replace(base, "~", env.Home(), 1)
	}

	if !filepath.IsAbs(base) {
		confdir := filepath.Dir(conffile)
		base = filepath.Join(confdir, base)
	}
	confMap[repoBase] = base

	return err
}

// Retruns the repository type.
func RepoType() types.RepoType {
	return types.RepoType(confMap[repoType])
}

// Returns the repository name.
func RepoName() string {
	return confMap[repoName]
}

// Returns the reopsitory base directory.
func RepoBase() string {
	return confMap[repoBase]
}

const confFilename = "config.yml"

func defaultConfFile() string {
	configHome := env.ConfigHome()
	configDir := filepath.Join(configHome, gonotes.Name)
	referenceConfigDir := filepath.Join(configHome, gonotes.ReferenceNotesName)

	// When the default config directory for `gonotes` is not exist,
	// try to refer the reference notes config directory.
	// If both are not exist, just use the default gonotes config directory.
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		if _, err := os.Stat(referenceConfigDir); !os.IsNotExist(err) {
			configDir = referenceConfigDir
		}
	}

	return filepath.Join(configDir, confFilename)
}

func loadConfYaml(filename string) map[string]string {
	file, _ := os.Open(filename)
	defer file.Close()

	keyPattern := regexp.MustCompile(`^(.+):\s+(.+)$`)
	quotedPattern := regexp.MustCompile(`^\"([^"]*)\"$`)
	confMap := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "---" {
			continue
		}

		m := keyPattern.FindStringSubmatch(line)
		key, value := m[1], m[2]

		// remove double quotation marks
		m = quotedPattern.FindStringSubmatch(value)
		if m != nil {
			value = m[1]
		}

		if key == repoType {
			value = string(confValueMap[value])
		}

		confMap[key] = value
	}
	return confMap
}
