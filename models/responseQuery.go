package models

type ResponseQuery struct {
	Description  string `json:"Description"`
	ResponseCode int32  `json:"ResponseCode"`
	Function     string `json:"Function"`
	Data         []Data `json:"Data"`
}
