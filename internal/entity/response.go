package entity

type SuccessFulResponse struct {
	Data    any  `json:"data"`
	Success bool `json:"success"`
}

type FailedResponse struct {
	ErrorCode string `json:"error_code"`
	ErrorData any    `json:"error_data"`
	Success   bool   `json:"success"`
}
