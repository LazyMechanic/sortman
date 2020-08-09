package questions

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/LazyMechanic/sortman/cli/dialog"
	"github.com/LazyMechanic/sortman/types"
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
	return selectOptions("What to do:", dialog.WhatToDo[0], dialog.WhatToDo)
}

func IsRequestCorrect() bool {
	return confirm("Is request correct:", true)
}

func Patterns(action types.Action) string {
	return input("Enter patterns:", "", fmt.Sprintf("Files which be %s to output directory. Separates patterns with %q, %q for example. Use '**/...' for recursive find, %q for example", action, ";", "*.png;*.jpg", "**/*.txt"), nil, survey.WithValidator(survey.Required))
}

func Exclude() string {
	return input("Enter exclusions:", "", fmt.Sprintf("Remove files or directories from selections. Separates exclusions with %q, %q for example", ";", "somefolder/;somefolder2/*.png"), nil)
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
	return absDir(outDefaultDir, input("Enter output directory:", outDefaultDir, "Set output directory for this request"))
}
