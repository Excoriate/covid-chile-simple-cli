package meta

import (
	"fmt"
	"github.com/urfave/cli"
)

func versionCommandImpl (){
	fmt.Println("Current CLI version: 1.0")
}

func LoadsCommands() []cli.Command {
	return  []cli.Command{
		{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "Check the version of the CLI",
			Action: func(c *cli.Context) {
				versionCommandImpl()
			},
		},
	}
}