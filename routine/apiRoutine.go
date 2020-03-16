package routine

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/matkinhig/go-topup/models"
	_ "github.com/mattn/go-oci8"
)

type apiRoutine struct {
	db *sqlx.DB
}

func NewAPIRoutine(db *sqlx.DB) *apiRoutine {
	return &apiRoutine{db}
}

func (a *apiRoutine) FindByCustID(dt *models.DataGet) (*models.ResponseGet, error) {
	// var err error
	// respGet := models.RequestGet{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		db, err := a.db.Begin()
		if err != nil {
			ch <- false
			return
		}
		sql := "select * from VB_DEPOSIT_AWARD where customer_id = ?"
		stmt, err := db.Prepare(sql)
		defer stmt.Close()
		if err != nil {
			ch <- false
			return
		}
		rows, err := stmt.Query(dt.CustomerID)
		if err != nil {
			ch <- false
			return
		}
		fmt.Println(rows)
	}(done)
	return nil, nil
}

// func (a *apiRoutine) FindAll() ([]models.DepositRepository, error) {
// var err error
// deposits := []models.DepositRepository{}
// done := make(chan bool)
// go func(ch chan<- bool) {
// 	defer close(ch)
// 	rows, err = a.db.Query("Select * from HALONG.VB_DEPOSIT_AWARD")
// 	if err != nil {
// 		ch <- false
// 		return
// 	}
// 	cols, err := rows.Columns()
// 	if err != nil {
// 		ch <- false
// 		return
// 	}
// 	store := []map[string]interface{}
// 	for rows.Next() {
// 		columns := make([]interface{}, len(cols))
// 		columnPointers := make([]interface{}, len(cols))
// 		for i, _ := range columns {
// 			columnPointers[i] = &columns[i]
// 		}
// 		if err := rows.Scan(columnPointers...); err != nil {
// 			ch <- false
// 			return
// 		}
// 		m := make(map[string]interface{})
// 		for i, colName := range cols {
// 			val := columnPointers[i].(*interface{})
// 			m[colName] = *val
// 		}
// 		store = append(store, m)
// 	}
// 	js, _ := json.Marshal(store)
// 	ch <- true
// }(done)
// if channels.OK(done) {
// 	return deposits, nil
// }
// return nil, err
// }
