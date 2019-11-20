package cli

import (
	"github.com/LazyMechanic/sortman/cli/commands"
	gocli "github.com/urfave/cli"
	"log"
	"os"
)

const (
	appName        string = "sortman"
	appUsage       string = "utility for sorting files by patterns to specific folders"
	appVersion     string = "0.1"
	appAuthorName  string = "LazyMechanic"
	appAuthorEmail string = "AsharnRus@gmail.com"
)

func Run() {
	process()
}

func process() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalln(r)
		}
	}()

	app := gocli.NewApp()
	app.Name = appName
	app.Version = appVersion
	app.Authors = []*gocli.Author{
		&gocli.Author{
			Name:  appAuthorName,
			Email: appAuthorEmail,
		},
	}
	app.Usage = appUsage

	app.Commands = []*gocli.Command{
		&commands.Copy,
		&commands.Move,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
