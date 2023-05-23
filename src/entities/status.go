package entities

type Status struct {
	Body `json:"status"`
}

type Body struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
