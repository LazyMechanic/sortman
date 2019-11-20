package commands

import (
	"github.com/LazyMechanic/sortman/cli/actions"
	gocli "github.com/urfave/cli"
)

var (
	Copy = gocli.Command{
		Name:      "copy",
		Aliases:   []string{"c"},
		Usage:     "copy files",
		ArgsUsage: "<from-directory> <to-directory>",
		Action:    actions.Copy,
	}

	Move = gocli.Command{
		Name:      "move",
		Aliases:   []string{"m"},
		Usage:     "move files",
		ArgsUsage: "<from-directory> <to-directory>",
		Action:    actions.Move,
	}
)
