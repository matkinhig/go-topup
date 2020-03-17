package routine

import (
	"errors"
	"fmt"

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

func (a *apiRoutine) FindByCustID(resGet *models.RequestGet) (*models.ResponseGet, error) {
	// var err error
	respGet := models.ResponseGet{}
	// respGet.Description = "Success"
	// respGet.ResponseCode = strconv.Itoa(http.StatusOK)
	respGet.Function = resGet.Function
	data := models.Data{}
	done := make(chan bool)
	custid := resGet.Data.CustomerID
	go func(ch chan<- bool) {
		defer close(ch)
		db, err := a.db.Beginx()
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
