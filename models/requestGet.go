package models

type RequestGet struct {
	AbstractModel
	Data []DataGet
}

type DataGet struct {
	CustomerID string `json:"customerid"`
}
