package models

import "errors"

type RequestPost struct {
	AbstractModel
	Data DataPost `json:"data"`
}

type DataPost struct {
	CustomerID       string  `json:"CustomerID"`
	EuserId          string  `json:"EuserId"`
	CustomerFullname string  `json:"CustomerFullname"`
	AccountDeposit   string  `json:"AccountDeposit"`
	Amount           float64 `json:"Amount"`
	Term             string  `json:"Term"`
	PhoneNumber      string  `json:"PhoneNumber"`
}

func (r *RequestPost) Validate() error {
	if r.AbstractModel.RequestID == "" {
		return errors.New("Required RequestID")
	}
	if r.AbstractModel.Function == "" {
		return errors.New("Required Function")
	}
	if r.AbstractModel.RequestDateTime == "" {
		return errors.New("Required RequestDateTime")
	}
	if r.Data.CustomerID == "" {
		return errors.New("Required CustomerID")
	}
	if r.Data.EuserId == "" {
		return errors.New("Required EuserId")
	}
	if r.Data.CustomerFullname == "" {
		return errors.New("Required CustomerFullname")
	}
	if r.Data.AccountDeposit == "" {
		return errors.New("Required AccountDeposit")
	}
	if r.Data.Amount == 0 {
		return errors.New("Required Amount")
	}
	if r.Data.Term == "" {
		return errors.New("Required Term")
	}
	if r.Data.PhoneNumber == "" {
		return errors.New("Required PhoneNumber")
	}
	return nil
}
