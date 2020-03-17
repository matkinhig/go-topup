package models

type ResponseCreate struct {
	Description  string `json:"Description"`
	ResponseCode int32  `json:"ResponseCode"`
	Function     string `json:"Function"`
	Data         Result `json:"Data"`
}
