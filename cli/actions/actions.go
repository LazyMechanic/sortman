package actions

import (
	"fmt"
	"github.com/LazyMechanic/sortman/internal/cli/dialog"
	"github.com/LazyMechanic/sortman/internal/cli/flags"
	"github.com/LazyMechanic/sortman/internal/cli/questions"
	"github.com/LazyMechanic/sortman/internal/execute"
	"github.com/LazyMechanic/sortman/internal/types"
	gocli "github.com/urfave/cli"
	"path/filepath"
	"strings"
)

var (
	config types.Config
)

func absDir(root string, path string) string {
	if root == path {
		return path
	}

	if filepath.IsAbs(path) {
		return path
	} else {
		out, err := filepath.Abs(filepath.Join(root, path))
		if err != nil {
			panic(err)
		}

		return out
	}
}

func toAsk() error {
	for {
		var whatToDo = questions.WhatToDo()
		switch whatToDo {
		case dialog.AddRequest:
			config.Requests = append(config.Requests, types.Request{
				Patterns:     strings.Split(questions.Patterns(), ";"),
				Exclude:      strings.Split(questions.Exclude(), ";"),
				InDirectory:  absDir(config.WorkingDirectory, questions.InDirectory(config.WorkingDirectory)),
				OutDirectory: absDir(flags.OutDirectory, questions.OutDirectory(flags.OutDirectory)),
			})
		case dialog.Execute:
			return &types.Execute{}
		case dialog.Cancel:
			return &types.Cancel{}
		}
	}

	return nil
}

func workingDirectory(c *gocli.Context) (string, error) {
	if c.NArg() > 2 {
		return "", fmt.Errorf("Invalid number of arguments")
	}

	if c.NArg() == 1 {
		return c.Args().Get(0), nil
	}

	return ".", nil
}

func execCommand(c *gocli.Context) error {
	var err error

	config.WorkingDirectory, err = workingDirectory(c)
	if err != nil {
		return err
	}

	err = toAsk()
	switch err.(type) {
	case *types.Execute:
		/* continue */
	case *types.Cancel:
		/* discard changes */
		return nil
	default:
		return err
	}

	err = execute.Execute(&config)

	return err
}

func Copy(c *gocli.Context) error {
	config.Action = dialog.CopyAction
	return execCommand(c)
}

func Move(c *gocli.Context) error {
	config.Action = dialog.MoveAction
	return execCommand(c)
}
