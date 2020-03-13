package models

import "time"

type Data struct {
	Euser_Id          string    `json:"Euser_Id"`
	Customer_Fullname string    `json:"Customer_Fullname"`
	Lottery_Code      string    `json:"Lottery_Code"`
	Account_Deposit   string    `json:"Account_Deposit"`
	Opened_Date       time.Time `json:"Opened_Date"`
	Closed_Date       time.Time `json:"Closed_Date"`
	Status            string    `json:"Status"`
	Amount            float64   `json:"Amount"`
	Term              string    `json:"Term"`
	Open_Award        int8      `json:"Open_Award"`
	Phone_Number      string    `json:"Phone_Number"`
}

type DepositRepository struct {
	AbstractModel
	Data []Data `json:"Data"`
}
