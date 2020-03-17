package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	_ "github.com/jmoiron/sqlx"
	"github.com/matkinhig/go-topup/database"
	"github.com/matkinhig/go-topup/models"
	"github.com/matkinhig/go-topup/repository"
	"github.com/matkinhig/go-topup/responses"
	"github.com/matkinhig/go-topup/routine"
	_ "github.com/mattn/go-oci8"
)

func DepositHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	requestGet := models.RequestGet{}
	err = json.Unmarshal(body, &requestGet)
	if err != nil {
		fmt.Println(err)
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = requestGet.Validate()
	if err != nil {
		fmt.Println(err)
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	thread := routine.NewAPIRoutine(db)
	func(getRepo repository.ApiRepository) {
		respGet, err := getRepo.FindByCustID(&requestGet)
		if err != nil {
			fmt.Println(err)
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		respGet.Description = "Success"
		respGet.ResponseCode = strconv.Itoa(http.StatusOK)
		responses.JSON(w, http.StatusOK, &respGet)
		return
	}(thread)
}
