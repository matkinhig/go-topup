package controllers

import (
	"fmt"
	"net/http"

	"github.com/matkinhig/go-topup/database"
	"github.com/matkinhig/go-topup/responses"
)

func DepositHandler(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	rows, err := db.Query("Select * from HALONG.VB_DEPOSIT_AWARD")
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	cols, err := rows.Columns()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
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
			responses.ERROR(w, http.StatusInternalServerError, err)
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
	fmt.Println(string(js))
}
