package models

import "errors"

type RequestUpdate struct {
	AbstractModel
	Data DataUpdate `json:"data"`
}

type DataUpdate struct {
	CustomerId     string `json:"CustomerId"`
	AccountDeposit string `json:"AccountDeposit"`
	UpdatedDate    string `json:"UpdatedDate"`
}

func (r *RequestUpdate) Validate() error {
	if r.AbstractModel.RequestID == "" {
		return errors.New("Required RequestID")
	}
	if r.AbstractModel.Function == "" {
		return errors.New("Required Function")
	}
	if r.AbstractModel.RequestDateTime == "" {
		return errors.New("Required RequestDateTime")
	}
	if r.Data.CustomerId == "" {
		return errors.New("Required CustomerID")
	}
	if r.Data.AccountDeposit == "" {
		return errors.New("Required AccountDeposit")
	}
	if r.Data.UpdatedDate == "" {
		return errors.New("Required UpdatedDate")
	}
	return nil
}
