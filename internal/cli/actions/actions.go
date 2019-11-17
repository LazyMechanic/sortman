package actions

import (
	"fmt"
	"github.com/LazyMechanic/sortman/internal/cli/dialog"
	"github.com/LazyMechanic/sortman/internal/cli/questions"
	"github.com/LazyMechanic/sortman/internal/types"
	gocli "github.com/urfave/cli"
)

const (
	copy types.Action = "copy"
	move types.Action = "move"
)

var (
	config types.Config
)

func toAsk() error {
	for {
		var whatToDo = questions.WhatToDo()
		switch whatToDo {
		case dialog.AddRequest:
			config.Requests = append(config.Requests, types.Request{
				Template:     questions.Template(),
				Exclude:      questions.Exclude(),
				InDirectory:  questions.InDirectory(),
				OutDirectory: questions.OutDirectory(),
			})
		case dialog.Execute:
			return &types.Execute{}
		case dialog.Cancel:
			return &types.Cancel{}
		}
	}

	return nil
}

func execute() error {
	// for _, request := range config.Requests {}
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

func Copy(c *gocli.Context) error {
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

	err = execute()

	return err
}

func Move(c *gocli.Context) error {
	var err error

	config.WorkingDirectory, err = workingDirectory(c)
	if err != nil {
		return err
	}

	return nil
}
