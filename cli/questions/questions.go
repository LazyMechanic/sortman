package questions

import (
	"fmt"
	"github.com/AlecAivazis/survey"
	"github.com/LazyMechanic/sortman/cli/dialog"
	"os"
	"path/filepath"
)

func confirm(msg string, defaultValue bool) bool {
	var answer bool
	var prompt = &survey.Confirm{
		Message: msg,
		Default: defaultValue,
	}

	err := survey.AskOne(prompt, &answer)
	if err != nil {
		panic(err)
	}
	return answer
}

func selectOptions(msg string, defaultValue string, options []string) string {
	var answer string
	var prompt = &survey.Select{
		Message: msg,
		Options: options,
		Default: defaultValue,
	}

	err := survey.AskOne(prompt, &answer)
	if err != nil {
		panic(err)
	}
	return answer
}

func input(msg string, defaultValue string, help string, opts ...survey.AskOpt) string {
	var answer string
	var prompt = &survey.Input{
		Message: msg,
		Default: defaultValue,
		Help:    help,
	}

	var err error

	err = survey.AskOne(prompt, &answer, opts...)
	if err != nil {
		panic(err)
	}

	return answer
}

func WhatToDo() string {
	var options = []string{
		dialog.AddRequest,
		dialog.Execute,
		dialog.Cancel,
	}
	return selectOptions("What to do:", options[0], options)
}

func IsRequestCorrect() bool {
	return confirm("Is request correct:", true)
}

func Patterns() string {
	return input("Enter patterns:", "", "Files which be copy or move to out dir. Patterns are listed with a ';', [*.png;*.jpg] for example", nil, survey.WithValidator(survey.Required))
}

func Exclude() string {
	return input("Enter exclusions:", "", "Files or directories to be dropped from the selection. Exclusions are listed with a ';', [somefolder/;somefolder2/*.png] for example", nil)
}

func isDirValidator(val interface{}) error {
	if val.(string) == "" {
		return nil
	}

	fileStat, err := os.Stat(val.(string))
	if err != nil {
		return err
	}

	if !fileStat.IsDir() {
		return fmt.Errorf("%q directory not found", fileStat.Name())
	}
	return nil
}

func absDir(root string, path string) string {
	if root == path {
		return path
	}

	if filepath.IsAbs(path) {
		return path
	}

	out, err := filepath.Abs(filepath.Join(root, path))
	if err != nil {
		panic(err)
	}

	return out
}

func InDirectory(inDefaultDir string) string {
	return absDir(inDefaultDir, input("Enter input directory:", inDefaultDir, "Set input directory for this request", survey.WithValidator(isDirValidator)))
}

func OutDirectory(outDefaultDir string) string {
	return absDir(outDefaultDir, input("Enter out directory:", outDefaultDir, "Set out directory for this request"))
}
