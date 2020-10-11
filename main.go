package main

import (
	"cli/v2/commands/covid"
	"cli/v2/commands/meta"
	"github.com/urfave/cli" // imports as package "cli"
	"log"
	"os"
)

var app = cli.NewApp()

func loadCLIInformation() {
	app.Name = "covidcli"
	app.Usage = "Simple CLI to fetch useful COVID information for Chilean people po weon!"
	app.Version = "1.0.0"
	app.Email = "alex@ideaup.cl"
	app.Author = "IdeaUP.cl"
}

func loadCLI(){
	loadCLIInformation()
	loadCLICommands()
	app.EnableBashCompletion = true
	app.Commands = loadCLICommands()
}

func loadCLICommands() []cli.Command {
	var commands []cli.Command
	
	commands = append(commands, meta.GetCommands()...)
	commands = append(commands, covid.GetCommands()...)
	return commands
}

func main() {
	loadCLI()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
