package questions

import (
	"github.com/AlecAivazis/survey"
	"github.com/LazyMechanic/sortman/internal/cli/dialog"
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

	err := survey.AskOne(prompt, &answer, opts...)
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

func Template() string {
	return input("Enter template:", "", "Files which be copy or move to out dir. Templates are listed with a space, [*.png *.jpg] for example", survey.WithValidator(survey.Required))
}

func Exclude() string {
	return input("Enter exclude:", "", "Files or directories to be dropped from the selection. Exclude are listed with a space, [somefolder/ somefolder2/*.png] for example")
}

func InDirectory() string {
	return input("Enter input directory:", "", "Replace current working directory for this request. Absolute or relative working directory path")
}

func OutDirectory() string {
	return input("Enter out directory:", "", "Set out directory for this request. Absolute or relative out directory path", survey.WithValidator(survey.Required))
}
