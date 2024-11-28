// gonotes/cmd/gonotes/main.go
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mnbi/gonotes"
	"github.com/mnbi/gonotes/internal/conf"
	"github.com/mnbi/gonotes/internal/textrepo"
)

var (
	versionFlag = flag.Bool("v", false, "show version")
	usageFlag   = flag.Bool("h", false, "show usage")
	configFile  string
)

func main() {
	flag.Usage = gonotes.Usage
	flag.StringVar(&configFile, "c", "", "specify a configuration file")

	flag.Parse()

	if *versionFlag {
		gonotes.ShowVersion()
	}

	if *usageFlag {
		flag.Usage()
	}

	args := flag.Args()

	fmt.Printf("ARGS: ")
	for _, arg := range args {
		fmt.Printf(" %s", arg)
	}
	fmt.Printf("\n")

	fmt.Printf("---- configuration settings ----\n")

	configFile, _ = filepath.Abs(configFile)
	if err := conf.Init(configFile); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("conf - repo type: %s\n", conf.RepoType())
	fmt.Printf("conf - repo name: %s\n", conf.RepoName())
	fmt.Printf("conf - repo base: %s\n", conf.RepoBase())

	fmt.Printf("---- repository settings ----\n")

	repo := textrepo.NewRepo()

	fmt.Printf("repo - type: %s\n", repo.Type)
	fmt.Printf("repo - name: %s\n", repo.Name)
	fmt.Printf("repo - root: %s\n", repo.Root)
	fmt.Printf("repo - ext: %s\n", repo.Ext)
}
