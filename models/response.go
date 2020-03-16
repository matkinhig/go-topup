package models

type Response struct {
	AbstractModel
	Description  string `json:"Description"`
	ResponseCode string `json:"ResponseCode"`
	Function     string `json:"Function"`
}

type ResponseGet struct {
	Description  string `json:"Description"`
	ResponseCode string `json:"ResponseCode"`
	Function     string `json:"Function"`
	Data         []Data
}
