package data

type stockData struct {
	Metadata   metaData             `json:"Meta Data"`
	Timeseries map[string]dailyData `json:"Time Series (Daily)"`
	OutputData outputData           `json:"Output Data"`
}

type metaData struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	OutputSize    string `json:"4. Output Size"`
	TimeZone      string `json:"5. Time Zone"`
}

type dailyData struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}

type outputData struct {
	NDays    int    `json:"N Days"`
	AvgClose string `json:"Average Close"`
}
