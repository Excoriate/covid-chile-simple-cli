package covid

/*
 All the commands should share the same structure:
	1. Command configuration
	2. Command handler
	3. Command implementation (called by the handler)
	4. Command or subcommands documentation, using heredoc
*/

import (
	"cli/v2/libs"
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

var API_URL_COVID_NATIONAL = "https://chile-coronapi1.p.rapidapi.com/v3/latest/nation"

func statsCmdHandler(c *cli.Context) {
	statsCmdImpl()
}

func helloCmdDoc() string {
	return heredoc.Doc(`
		covidcli stats

		EXAMPLES:
		    covidcli stats
			covidcli stats -region santiago`)
}

func statsCmdImpl() {
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)  // Build our new spinner
	s.Color("bgBlack", "bold", "fgRed")
	s.Start()
	
	// Building HTTP request arguments
	var httpRequest  libs.HTTPRequestArgs
	var httpHeaders []libs.HTTPHeaders
	
	var httpHeaderHost libs.HTTPHeaders
	httpHeaderHost.Key = "x-rapidapi-host"
	httpHeaderHost.Value = "chile-coronapi1.p.rapidapi.com"
	
	var httpHeaderKey libs.HTTPHeaders
	httpHeaderKey.Key = "x-rapidapi-key"
	httpHeaderKey.Value = "f255b0bc76msh76c5f8695aae921p10ccfbjsn2dc36c8dabde"
	
	httpHeaders = append(httpHeaders, httpHeaderHost)
	httpHeaders = append(httpHeaders, httpHeaderKey)
	
	httpRequest.Uri = API_URL_COVID_NATIONAL
	httpRequest.Headers = httpHeaders

	response := libs.ExecHttpRequest(httpRequest)
	var nationalCovidResultStruct NationalResult
	err :=  json.Unmarshal([]byte(response), &nationalCovidResultStruct)
	
	if err != nil {
		log.Fatal("Error. Could not parse JSON response")
	}
	s.Stop()
	
	// TODO: split this logic in a separate function
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Day", "Confirmed Cases", "Mortality"})
	
	data := [][]string{
		[]string{nationalCovidResultStruct.DayReported, strconv.Itoa(int(nationalCovidResultStruct.Confirmed)), strconv.Itoa(int(nationalCovidResultStruct.DeathsReportedTotal))},
		[]string{nationalCovidResultStruct.DayReported, strconv.FormatFloat(nationalCovidResultStruct.ConfirmedPer100k, 'f', 6, 64) + " per 100K", strconv.FormatFloat(nationalCovidResultStruct.DeathsReportedPer100k, 'f', 6, 64)  + " per 100K"},
		[]string{nationalCovidResultStruct.DayReported, strconv.FormatFloat(nationalCovidResultStruct.ConfirmedPerMillion, 'f', 6, 64)  + " per 1M", strconv.FormatFloat(nationalCovidResultStruct.DeathsReportedPerMillion, 'f', 6, 64)  + " per 1 M"},
	}
	
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}

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
