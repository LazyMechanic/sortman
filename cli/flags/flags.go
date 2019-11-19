package flags

import (
	gocli "github.com/urfave/cli"
)

var (
	Recursive    bool
	OutDirectory string
)

var (
	RecursiveFlag gocli.Flag = &gocli.BoolFlag{
		Name:        "r, recursive",
		Usage:       "use recursive search",
		Destination: &Recursive,
		Value:       false,
	}

	OutDirectoryFlag gocli.Flag = &gocli.StringFlag{
		Name:        "o, out-dir",
		Usage:       "set root out directory",
		Destination: &OutDirectory,
		Value:       "",
	}
)
