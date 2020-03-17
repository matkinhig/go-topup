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
		db, err := a.db.Beginx()
		defer a.db.Close()
		if err != nil {
			ch <- false
			return
		}
		stmt, err := db.Preparex(`select * from HALONG.VB_DEPOSIT_AWARD where customer_id= :custid`)
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
	return &respGet, nil
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
		data.Phone_Number = resPost.Data.CustomerFullname
		data.Customer_Id = resPost.Data.CustomerID
		db := a.db.MustBegin()
		defer a.db.Close()
		if err != nil {
			ch <- false
			return
		}
		sql := `INSERT INTO HALONG.VB_DEPOSIT_AWARD 
		(Euser_Id,Customer_Fullname,Lottery_Code,Account_Deposit,Opened_Date,Status,Amount,Term,Open_Award,Phone_Number,customer_id) VALUES
		(:Euser_Id, :Customer_Fullname, :Lottery_Code, :Account_Deposit, :Opened_Date, :Status, :Amount, :Term, :Open_Award, :Phone_Number, :Customer_Id)`
		rs, err := db.Exec(sql,
			data.Euser_Id, data.Customer_Fullname, data.Lottery_Code,
			data.Account_Deposit, data.Opened_Date, data.Status,
			data.Amount, data.Term, data.Open_Award, data.Phone_Number, data.Customer_Id)
		fmt.Println(err)
		if err != nil {
			ch <- false
			return
		}
		err = db.Commit()
		if err != nil {
			ch <- false
			return
		}
		_, err = rs.RowsAffected()
		if err != nil {
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
