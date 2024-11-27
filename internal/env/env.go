// gonotes/internal/env/env.go

// The package to provide the access point to retrieve environment
// variables, like "HOME", or "XDG_CONFIG_HOME".
//
// All interfaces won't return any errors. Instead, they will return a
// default value, when a target variable is not set.
package env

import (
	"os"
	"path/filepath"
)

type envStat struct {
	value string
	stat  bool
}

var keys = [...]string{
	"HOME",
	"XDG_CONFIG_HOME",
}

var Env map[string]envStat = make(map[string]envStat)

func init() {
	for _, key := range keys {
		val, ok := os.LookupEnv(key)
		Env[key] = envStat{value: val, stat: ok}
	}
}

// Returns the value of "HOME". Its the default value is "/".
func Home() string {
	home, err := os.UserHomeDir()
	if err != nil {
		home = "/"
	}
	return home
}

// Returns the value of "XDG_CONFIG_HOME". When it is not set, use the
// return value of os.UserConfigDir(). If the call returns an error,
// the default value for the config directory is used, which is
// ".config" in the user home directory.
func ConfigHome() string {
	var confdir string
	var err error

	if Env["XDG_CONFIG_HOME"].stat {
		confdir = Env["XDG_CONFIG_HOME"].value
	} else {
		// XDG_CONFIG_HOME is not set
		confdir, err = os.UserConfigDir()
		if err != nil {
			confdir = defaultConfigDir()
		}
	}

	return confdir
}

func defaultConfigDir() string {
	return filepath.Join(Home(), ".config")
}
