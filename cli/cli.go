package cli

import (
	"github.com/LazyMechanic/sortman/internal/cli/commands"
	gocli "github.com/urfave/cli"
	"log"
	"os"
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
	app.Name = "sortman"
	app.Version = "0.1"
	app.Authors = []*gocli.Author{
		&gocli.Author{
			Name:  "LazyMechanic",
			Email: "asharnrus@gmail.com",
		},
	}
	app.Usage = "utility for sorting files by patterns to specific folders"

	app.Commands = []*gocli.Command{
		&commands.Copy,
		&commands.Move,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
