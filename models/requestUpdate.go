package models

type RequestUpdate struct {
	AbstractModel
	Data []DataUpdate `json:"data"`
}

type DataUpdate struct {
	CustomerId  string `json:"customerid"`
	AwardCode   string `json:"awardcode"`
	UpdatedDate string `json:"updateddate"`
}
