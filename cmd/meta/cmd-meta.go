package meta

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/urfave/cli"
)

func versionCommandImpl (){
	fmt.Println("Current CLI version: 1.0")
}

func versionHereDoc() string {
	return heredoc.Doc(`
		covidcli version [command options] <key arn or alias>

		EXAMPLES:
		   covidcli -version`)
}

func GetCommands() []cli.Command {
	return  []cli.Command{
		{
			Name:    "version",
			Aliases: []string{"v"},
			UsageText: versionHereDoc(),
			Usage:   "Check the version of the CLI",
			Action: func(c *cli.Context) {
				versionCommandImpl()
			},
		},
	}
}