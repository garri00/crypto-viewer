package entities

type ExchangeRate struct {
	Quotes struct {
		USDUAH float64 `json:"USDUAH"`
	} `json:"quotes"`
	Source    string `json:"source"`
	Success   bool   `json:"success"`
	Timestamp int    `json:"timestamp"`
}
