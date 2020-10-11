package covid

/*
 All the commands should share the same structure:
	1. Command configuration
	2. Command handler
	3. Command implementation (called by the handler)
	4. Command or subcommands documentation, using heredoc
*/

import (
	"cli/v2/utils"
	"encoding/json"
	"github.com/MakeNowJust/heredoc"
	"github.com/briandowns/spinner"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
	"log"
	"os"
	"strconv"
	"time"
)

func statsCmdHandler(c *cli.Context) {
	statsCmdImpl()
}

func helloCmdDoc() string {
	return heredoc.Doc(`
		covidcli hello

		EXAMPLES:
		   covidcli -hello`)
}

func statsCmdImpl() {
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)  // Build our new spinner
	s.Color("bgBlack", "bold", "fgRed")
	s.Start()
	
	var url = "https://chile-coronapi1.p.rapidapi.com/v3/latest/nation"
	response := utils.ExecHttpRequest(url, "GET")
	var nationalCovidResultStruct NationalResult
	err :=  json.Unmarshal([]byte(response), &nationalCovidResultStruct)
	
	if err != nil {
		log.Fatal("Error. Could not parse JSON response")
	}
	s.Stop()
	
	// TODO: split this logic in a separate function
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Category", "KPI"})
	
	data := [][]string{
		[]string{"Confirmed", strconv.Itoa(int(nationalCovidResultStruct.Confirmed))},
		[]string{"Confirmed per 100 K", strconv.FormatFloat(nationalCovidResultStruct.ConfirmedPer100k, 'f', 6, 64)},
	}
	
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
	
	
	// return nationalCovidResultStruct
}

/*
func getCmdFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "test",
			Usage: "Test flag",
			Value: "",
		},
	}
}

 */

func GetCommands() []cli.Command {
	return []cli.Command{
		{
			Name:      "stats",
			Aliases:   []string{"s"},
			UsageText: helloCmdDoc(), // TODO: Update command documentation
			Usage:     "Get specific CoronaVirus statistics in Chile",
			Action:    statsCmdHandler,
		},
	}
}
