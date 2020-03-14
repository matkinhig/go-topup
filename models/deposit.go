package models

import (
	"time"
)

type Nulltime struct {
	Time  time.Time
	Valid bool
}

func (nt *Nulltime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

type Data struct {
	Euser_Id          string    `db:"EUSER_ID" json:"Euser_Id"`
	Customer_Fullname string    `db:"CUSTOMER_FULLNAME" json:"Customer_Fullname"`
	Lottery_Code      string    `db:"LOTTERY_CODE" json:"Lottery_Code"`
	Account_Deposit   string    `db:"ACCOUNT_DEPOSIT" json:"Account_Deposit"`
	Opened_Date       time.Time `db:"OPENED_DATE" json:"Opened_Date"`
	Closed_Date       Nulltime  `db:"CLOSED_DATE" json:"Closed_Date"`
	Status            string    `db:"STATUS" json:"Status"`
	Amount            float64   `db:"AMOUNT" json:"Amount"`
	Term              string    `db:"TERM" json:"Term"`
	Open_Award        int8      `db:"OPEN_AWARD" json:"Open_Award"`
	Phone_Number      string    `db:"PHONE_NUMBER" json:"Phone_Number"`
}

type DepositRepository struct {
	AbstractModel
	Data []Data `json:"Data"`
}
