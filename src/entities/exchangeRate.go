package entities

type ExchangeRate struct {
	Quotes    `json:"quotes"`
	Source    string `json:"source"`
	Success   bool   `json:"success"`
	Timestamp int    `json:"timestamp"`
}

type Quotes struct {
	USDUAH float64 `json:"USDUAH"`
}
