package covid

import (
	"fmt"
	"github.com/urfave/cli"
)

func LoadsCommands() []cli.Command {
	return  []cli.Command{
		{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "Check the version of the CLI",
			Action: func(c *cli.Context) {
				fmt.Println("something")
			},
		},
	}
}