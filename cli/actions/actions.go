package actions

import (
	"fmt"
	"github.com/LazyMechanic/sortman/cli/dialog"
	"github.com/LazyMechanic/sortman/cli/questions"
	"github.com/LazyMechanic/sortman/execute"
	"github.com/LazyMechanic/sortman/types"
	gocli "github.com/urfave/cli"
	"os"
	"path/filepath"
	"strings"
)

func removeEmptyStrings(strs []string) []string {
	var result []string

	for _, str := range strs {
		if str != "" {
			result = append(result, str)
		}
	}

	return result
}

func showRequests(requests *[]types.Request) {
	if len(*requests) == 0 {
		fmt.Println("No have requests")
	}

	for index, _ := range *requests {
		(*requests)[index].Fprint(os.Stdout)
	}
}

func toAsk(config *types.Config) error {
	for {
		var whatToDo = questions.WhatToDo()
		switch whatToDo {
		case dialog.AddRequest:
			config.Requests = append(config.Requests, types.Request{
				Patterns:     removeEmptyStrings(strings.Split(questions.Patterns(config.Action), ";")),
				Exclusions:   removeEmptyStrings(strings.Split(questions.Exclude(), ";")),
				InDirectory:  questions.InDirectory(config.InDirectory),
				OutDirectory: questions.OutDirectory(config.OutDirectory),
			})

			if !questions.IsRequestCorrect() {
				continue
			}
		case dialog.ShowRequests:
			showRequests(&config.Requests)
		case dialog.Execute:
			return &types.Execute{}
		case dialog.Cancel:
			return &types.Cancel{}
		}
	}

	return nil
}

func inDirectory(c *gocli.Context) (string, error) {
	if c.NArg() > 2 {
		return "", fmt.Errorf("Invalid number of arguments")
	}

	if c.NArg() > 0 {
		return filepath.Abs(c.Args().Get(0))
	}

	return filepath.Abs(".")
}

func outDirectory(c *gocli.Context) (string, error) {
	if c.NArg() > 2 {
		return "", fmt.Errorf("Invalid number of arguments")
	}

	if c.NArg() == 2 {
		return filepath.Abs(c.Args().Get(1))
	}

	return filepath.Abs(".")
}

func execCommand(c *gocli.Context, config *types.Config) error {
	var err error

	if c.NArg() > 2 {
		return fmt.Errorf("Invalid number of arguments")
	}

	config.InDirectory, err = inDirectory(c)
	if err != nil {
		return err
	}

	config.OutDirectory, err = outDirectory(c)
	if err != nil {
		return err
	}

	err = toAsk(config)
	switch err.(type) {
	case *types.Execute:
		/* continue */
	case *types.Cancel:
		/* discard changes */
		return nil
	default:
		return err
	}

	err = execute.Execute(config)

	return err
}

func Copy(c *gocli.Context) error {
	var config types.Config
	config.Action = dialog.CopyAction

	return execCommand(c, &config)
}

func Move(c *gocli.Context) error {
	var config types.Config
	config.Action = dialog.MoveAction

	return execCommand(c, &config)
}
