package models

import "errors"

type RequestGet struct {
	AbstractModel
	Data DataGet
}

type DataGet struct {
	CustomerID string `json:"customerid"`
}

func (r *RequestGet) Validate() error {
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
	return nil
}
