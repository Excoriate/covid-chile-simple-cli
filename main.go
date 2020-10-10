package covid_chile_simple_cli

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name: "ideaupCLI",
		Usage: "Use this CLI to do common operations like: create a new service based on a template, etc.",
		Action: func(c *cli.Context) error {
			fmt.Println("Chan chan! testing my CLI")
			return nil
		},
	}
	
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
