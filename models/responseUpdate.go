package models

type ResponseUpdate struct {
	Description  string `json:"Description"`
	ResponseCode int32  `json:"ResponseCode"`
	Function     string `json:"Function"`
	Data         Result `json:"Data"`
}
