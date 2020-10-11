package covid

type NationalResult struct {
	Confirmed    int32    `json:"confirmed"`
	ConfirmedPer100k        float64    `json:"confirmed_per_100k"`
	ConfirmedPerMillion     float64 `json:"confirmed_per_million"`
	DayReported string   `json:"day"`
	DeathsReportedTotal int32   `json:"deaths"`
	DeathsReportedPer100k float64   `json:"deaths_per_100k"`
	DeathsReportedPerMillion float64   `json:"deaths_per_million"`
}