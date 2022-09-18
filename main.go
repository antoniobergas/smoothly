package main

import (
	"flag"
)

func main() {
	flag.Parse()

	if install {
		runInstall()
	}

	if version {
		runVersion()
	}

	if run {
		runSolution()
	}

	if clean {
		runClean()
	}
}
