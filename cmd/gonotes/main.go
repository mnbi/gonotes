// gonotes/cmd/gonotes/main.go
package main

import (
	"flag"
	"fmt"

	"github.com/mnbi/gonotes"
)

var (
	versionFlag = flag.Bool("v", false, "show version")
	usageFlag   = flag.Bool("h", false, "show usage")
)

func main() {
	flag.Usage = gonotes.Usage
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
}
