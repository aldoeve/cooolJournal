package setup

import (
    "flag"
)

func SetupFlags(flags map[string]bool) {
	reset := flag.Bool("resetdb", false, "reset database")

	flag.Parse()

	if *reset {
		flags["resetdb"] = true
	}
}