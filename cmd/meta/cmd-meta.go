package meta

/*
 All the commands should share the same structure:
	1. Command configuration
	2. Command handler
	3. Command implementation (called by the handler)
	4. Command or subcommands documentation, using heredoc
 */

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/urfave/cli"
)

func helloCmdHandler(c *cli.Context) {
	helloCmdImpl()
}

func helloCmdDoc() string {
	return heredoc.Doc(`
		covidcli hello

		EXAMPLES:
		   covidcli -hello`)
}

func helloCmdImpl() string {
	return"Current CLI version: 1.0, hello motherfucker!"
}

func GetCommands() []cli.Command {
	return  []cli.Command{
		{
			Name:      "hello",
			Aliases:   []string{"h"},
			UsageText: helloCmdDoc(),
			Usage:     "Just for validate that the CLI was successfully installed :)",
			Action:    helloCmdHandler,
		},
	}
}