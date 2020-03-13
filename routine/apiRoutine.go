package routine

import (
	"database/sql"

	"github.com/matkinhig/go-topup/channels"
	"github.com/matkinhig/go-topup/models"
)

type apiRoutine struct {
	db *sql.DB
}

func NewAPIRoutine(db *sql.DB) *apiRoutine {
	return &apiRoutine{db}
}

func (a *apiRoutine) FindAll() ([]models.DepositRepository, error) {
	var err error
	deposits := []models.DepositRepository{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rows, err = a.db.Query("Select * from HALONG.VB_DEPOSIT_AWARD")
		if err != nil {
			ch <- false
			return
		}
		cols, err := rows.Columns()
		if err != nil {
			ch <- false
			return
		}
		store := []map[string]interface{}
		for rows.Next() {
			columns := make([]interface{}, len(cols))
			columnPointers := make([]interface{}, len(cols))
			for i, _ := range columns {
				columnPointers[i] = &columns[i]
			}
			if err := rows.Scan(columnPointers...); err != nil {
				ch <- false
				return
			}
			m := make(map[string]interface{})
			for i, colName := range cols {
				val := columnPointers[i].(*interface{})
				m[colName] = *val
			}
			store = append(store, m) 
		}
		js, _ := json.Marshal(store)
		ch <- true
	}(done)
	if channels.OK(done) {
		return deposits, nil
	}
	return nil, err
}
