package actions

import (
	"fmt"
	"github.com/LazyMechanic/sortman/internal/cli/dialog"
	"github.com/LazyMechanic/sortman/internal/cli/flags"
	"github.com/LazyMechanic/sortman/internal/cli/questions"
	"github.com/LazyMechanic/sortman/internal/types"
	gocli "github.com/urfave/cli"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	copy types.Action = "copy"
	move types.Action = "move"
)

var (
	config types.Config
)

func absDir(root string, path string) string {
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
				InDirectory:  absDir(config.WorkingDirectory, questions.InDirectory()),
				OutDirectory: absDir(flags.OutDirectory, questions.OutDirectory()),
			})
		case dialog.Execute:
			return &types.Execute{}
		case dialog.Cancel:
			return &types.Cancel{}
		}
	}

	return nil
}

func filesToExecute(request types.Request) ([]string, error) {
	if flags.Recursive {
		var err = filepath.Walk(request.InDirectory, func(path string, info os.FileInfo, err error) error {

			return nil
		})
		if err != nil {
			return []string{}, err
		}
	} else {
		files, err := ioutil.ReadDir(request.InDirectory)
		if err != nil {
			return []string{}, err
		}

		for _, file := range files {
			if file.IsDir() {
				continue
			}

			/*for _, pattern := range request.Patterns {

			}*/
		}
	}

	return []string{}, nil
}

func execute() error {
	/*
	var err error
	for _, request := range config.Requests {
		files, err := filesToExecute(request)
		if err != nil {
			return err
		}
	}
	*/
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

	err = execute()

	return err
}

func Copy(c *gocli.Context) error {
	config.Action = copy
	return execCommand(c)
}

func Move(c *gocli.Context) error {
	config.Action = move
	return execCommand(c)
}
