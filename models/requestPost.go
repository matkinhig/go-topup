package models

type RequestPost struct {
	AbstractModel
	Data []DataPost `json:"data"`
}

type DataPost struct {
	CustomerID       string  `json:"CustomerID"`
	EuserId          string  `json:"EuserId"`
	CustomerFullname string  `json:"CustomerFullname"`
	AccountDeposit   string  `json:"AccountDeposit"`
	Amount           float64 `json:"Amount"`
	Term             string  `json:"Term"`
	PhoneNumber      string  `json:"Phone_Number"`
}
