package commands

import (
	"github.com/LazyMechanic/sortman/internal/cli/actions"
	gocli "github.com/urfave/cli"
)

var (
	Copy = gocli.Command{
		Name:      "copy",
		Aliases:   []string{"c"},
		Usage:     "copy files",
		ArgsUsage: "<working-directory>",
		Flags: []gocli.Flag{
		},
		Action: actions.Copy,
	}

	Move = gocli.Command{
		Name:      "move",
		Aliases:   []string{"m"},
		Usage:     "move files",
		ArgsUsage: "<working-directory>",
		Flags: []gocli.Flag{
		},
		Action: actions.Move,
	}
)