package controllers

import (
	"fmt"
	"net/http"

	_ "github.com/jmoiron/sqlx"
	"github.com/matkinhig/go-topup/database"
	"github.com/matkinhig/go-topup/models"
	"github.com/matkinhig/go-topup/responses"
	_ "github.com/mattn/go-oci8"
)

func DepositHandler(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	defer db.Close()
	sql := "Select * from HALONG.VB_DEPOSIT_AWARD"
	rows, err := db.Queryx(sql)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	data := models.Data{}
	for rows.Next() {
		err := rows.StructScan(&data)
		if err != nil {
			fmt.Println(err)
			responses.ERROR(w, http.StatusPartialContent, err)
			return
		}
		fmt.Println(data)
	}
	// cols, err := rows.Columns()
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }
	// fmt.Printf("%t", reflect.TypeOf(cols))
	// fmt.Println(cols)
	// var store []map[string]interface{}
	// for rows.Next() {
	// 	columns := make([]interface{}, len(cols))
	// 	columnsPointer := make([]interface{}, len(cols))
	// 	for i, _ := range columns {
	// 		columnsPointer[i] = &columns[i]
	// 	}
	// 	if err = rows.Scan(columnsPointer...); err != nil {
	// 		responses.ERROR(w, http.StatusPartialContent, err)
	// 		return
	// 	}
	// 	m := make(map[string]interface{})
	// 	for i, colName := range cols {
	// 		val := columnsPointer[i].(*interface{})
	// 		m[colName] = *val
	// 	}
	// 	store = append(store, m)
	// }
	// // data := []models.Data{}
	// // json.Unmarshal(store, &data)
	// js, _ := json.Marshal(store)
	// fmt.Println(js)
	// db, err := database.Connect()
	// if err != nil {
	// 	responses.ERROR(w, http.StatusInternalServerError, err)
	// 	return
	// }
	// defer db.Close()
	// rows, err := db.Query("Select * from HALONG.VB_DEPOSIT_AWARD")
	// if err != nil {
	// 	responses.ERROR(w, http.StatusInternalServerError, err)
	// 	return
	// }
	// cols, err := rows.Columns()
	// if err != nil {
	// 	responses.ERROR(w, http.StatusInternalServerError, err)
	// 	return
	// }
	// fmt.Println(cols)
	// store := []map[string]interface{}
	// for rows.Next() {
	// 	columns := make([]interface{}, len(cols))
	// 	columnPointers := make([]interface{}, len(cols))
	// 	for i, _ := range columns {
	// 		columnPointers[i] = &columns[i]
	// 	}
	// 	if err := rows.Scan(columnPointers...); err != nil {
	// 		responses.ERROR(w, http.StatusInternalServerError, err)
	// 		return
	// 	}
	// 	m := make(map[string]interface{})
	// 	for i, colName := range cols {
	// 		val := columnPointers[i].(*interface{})
	// 		m[colName] = *val
	// 	}
	// 	store = append(store, m)
	// }
	// js, _ := json.Marshal(store)
	// fmt.Println(string(js))

	// db, err := sql.Open(config.DBDRIVER, config.DBURL)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// println("Connection succcess!!")
	// rows, err := db.Query("SELECT sysdate FROM dual")
	// if err != nil {
	// 	log.Fatalln("err:", err.Error)
	// }
	// var (
	// 	sysdate string
	// )
	// for rows.Next() {
	// 	if err = rows.Scan(&sysdate); err != nil {
	// 		log.Fatalln("error fetching", err)
	// 	}
	// 	log.Println(sysdate)
	// }
}
