package gonotes

import (
	"flag"
	"fmt"
	"os"
)

const name = "gonotes"
const description = "Gonotes - a simple utility to write a note"

var version = "LATEST_TAG"
var revision = "HEAD"

func Usage() {
	fmt.Fprintf(os.Stderr, "%s\n", description)
	fmt.Fprintf(os.Stderr, "usage: %s [options] [file]\n", name)
	flag.PrintDefaults()
	os.Exit(2)
}

func ShowVersion() {
	fmt.Fprintf(os.Stderr, "%s version %s (rev:%s)\n", name, version, revision)
	os.Exit(2)
}
