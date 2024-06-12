package corpus

type Data struct {
	TIPLOCData []struct {
		NLC        int    `json:"NLC"`
		STANOX     string `json:"STANOX"`
		TIPLOC     string `json:"TIPLOC"`
		ThreeAlpha string `json:"3ALPHA"`
		UIC        string `json:"UIC"`
		NLCDesc    string `json:"NLCDESC"`
		NLCDesc16  string `json:"NLCDESC16"`
	} `json:"TIPLOCDATA"`
}
