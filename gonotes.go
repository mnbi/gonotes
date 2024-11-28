package gonotes

import (
	"flag"
	"fmt"
	"os"
)

const Name = "gonotes"
const description = "Gonotes - a simple utility to write a note"
const ReferenceNotesName = "rbnotes"

var version = "LATEST_TAG"
var revision = "HEAD"

func Usage() {
	fmt.Fprintf(os.Stderr, "%s\n", description)
	fmt.Fprintf(os.Stderr, "usage: %s [options] [file]\n", Name)
	flag.PrintDefaults()
	os.Exit(2)
}

func ShowVersion() {
	fmt.Fprintf(os.Stderr, "%s version %s (rev:%s)\n", Name, version, revision)
	os.Exit(2)
}
