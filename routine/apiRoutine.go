package routine

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/matkinhig/go-topup/channels"
	"github.com/matkinhig/go-topup/models"
	_ "github.com/mattn/go-oci8"
)

type apiRoutine struct {
	db *sqlx.DB
}

func NewAPIRoutine(db *sqlx.DB) *apiRoutine {
	return &apiRoutine{db}
}

func (a *apiRoutine) FindByCustID(resGet *models.RequestGet) (*models.ResponseQuery, error) {
	respGet := models.ResponseQuery{}
	respGet.Function = resGet.Function
	data := models.Data{}
	done := make(chan bool)
	custid := resGet.Data.CustomerID
	go func(ch chan<- bool) {
		defer close(ch)
		defer a.db.Close()
		stmt, err := a.db.Preparex(`select * from HALONG.VB_DEPOSIT_AWARD where customer_id= :custid`)
		defer stmt.Close()
		if err != nil {
			ch <- false
			return
		}
		rows, err := stmt.Queryx(custid)
		if err != nil {
			ch <- false
			return
		}
		for rows.Next() {
			err = rows.StructScan(&data)
			if err != nil {
				ch <- false
				return
			}
			respGet.Data = append(respGet.Data, data)
			fmt.Println(respGet)
			fmt.Println(respGet.Data)
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return &respGet, nil
	}
	if len(respGet.Data) < 1 {
		return &respGet, errors.New("Cant Find Any Customer")
	}
	return nil, errors.New("Unprocess")
}

func (a *apiRoutine) CreateDeposit(resPost *models.RequestPost) (*models.ResponseCreate, error) {
	var err error
	respPost := models.ResponseCreate{}
	respPost.Function = resPost.Function
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		data := models.Data{}
		num, err := strconv.Atoi(resPost.Data.AccountDeposit)
		if err != nil {
			ch <- false
			return
		}
		data.Lottery_Code = strconv.Itoa(num)
		data.Euser_Id = resPost.Data.EuserId
		data.Customer_Fullname = resPost.Data.CustomerFullname
		data.Account_Deposit = resPost.Data.AccountDeposit
		data.Opened_Date = time.Now()
		data.Status = "1"
		data.Amount = resPost.Data.Amount
		data.Term = resPost.Data.Term
		data.Open_Award = 1
		data.Phone_Number = resPost.Data.PhoneNumber
		data.Customer_Id = resPost.Data.CustomerID
		db := a.db
		defer db.Close()
		sql := `INSERT INTO HALONG.VB_DEPOSIT_AWARD 
		(Euser_Id,Customer_Fullname,Lottery_Code,Account_Deposit,Opened_Date,Status,Amount,Term,Open_Award,Phone_Number,customer_id) VALUES
		(:Euser_Id, :Customer_Fullname, :Lottery_Code, :Account_Deposit, :Opened_Date, :Status, :Amount, :Term, :Open_Award, :Phone_Number, :Customer_Id)`
		rs, err := db.Exec(sql,
			data.Euser_Id, data.Customer_Fullname, data.Lottery_Code,
			data.Account_Deposit, data.Opened_Date, data.Status,
			data.Amount, data.Term, data.Open_Award, data.Phone_Number, data.Customer_Id)
		if err != nil {
			ch <- false
			return
		}
		i, err := rs.RowsAffected()
		if err != nil {
			ch <- false
			return
		}
		if i < 1 {
			err = errors.New("Cant Create The Account Deposit, please try again later")
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return &respPost, nil
	}
	return nil, err
}

func (a *apiRoutine) UpdateDeposit(resUpdate *models.RequestUpdate) (*models.ResponseUpdate, error) {
	var err error
	respUpdate := models.ResponseUpdate{}
	respUpdate.Function = resUpdate.Function
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		db := a.db
		defer db.Close()
		sql := `UPDATE HALONG.VB_DEPOSIT_AWARD SET CLOSED_DATE=:tm, AMOUNT=:am, OPEN_AWARD=:ow WHERE CUSTOMER_ID=:custid and ACCOUNT_DEPOSIT=:accde`
		rs, err := db.Exec(sql,
			time.Now(), 0, 0, resUpdate.Data.CustomerId, resUpdate.Data.AccountDeposit)
		if err != nil {
			ch <- false
			return
		}
		i, err := rs.RowsAffected()
		if err != nil {
			ch <- false
			return
		}
		if i < 1 {
			err = errors.New("Cant Update The Account Deposit, please try again later")
			ch <- false
			return
		}
		fmt.Println(i)
		ch <- true
	}(done)
	if channels.OK(done) {
		return &respUpdate, nil
	}
	return nil, err
}
