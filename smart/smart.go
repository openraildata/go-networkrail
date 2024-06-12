package smart

// Data holds the SMART data.
type Data struct {
	BerthData []struct {
		// Train Describer area
		TD string `json:"TD"`
		// TD berth that movement is from
		FromBerth string `json:"FROMBERTH"`
		// TD berth that movement is to
		ToBerth string `json:"TOBERTH"`
		// Line which movement is from
		FromLine string `json:"FROMLINE"`
		// Line which movement is to
		ToLine string `json:"TOLINE"`
		// Difference between the time the berth event occurs and the time to be recorded in TRUST
		// in seconds
		BerthOffset BerthOffset `json:"BERTHOFFSET"`
		Platform    string      `json:"PLATFORM"`
		Event       Event       `json:"EVENT"`
		Route       string      `json:"ROUTE"`
		STANOX      string      `json:"STANOX"`
		STANME      string      `json:"STANME"`
		StepType    StepType    `json:"STEPTYPE"`
		Comment     string      `json:"COMMENT"`
	} `json:"BERTHDATA"`
}
