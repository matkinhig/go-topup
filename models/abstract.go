package models

type AbstractModel struct {
	RequestID       string `json:"RequestID"`
	RequestDateTime string `json:"RequestDateTime"`
	Function        string `json:"Function"`
}
