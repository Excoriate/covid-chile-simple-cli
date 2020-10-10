package main

import (
	"cli/v2/cmd/meta"
	"github.com/urfave/cli" // imports as package "cli"
	"log"
	"os"
)

var app = cli.NewApp()

func cliInfo() {
	app.Name = "covidcli"
	app.Usage = "Simple CLI to fetch useful COVID information for Chilean people po weon!"
	app.Version = "1.0.0"
	app.Email = "alex@ideaup.cl"
	app.Author = "IdeaUP.cl"
}

func loadMetaCommands(){
	app.Commands = meta.LoadsCommands()
}

func main() {
	cliInfo()
	loadMetaCommands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
