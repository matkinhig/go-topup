package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

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
	resquest := models.AbstractModel{}
	err = json.Unmarshal(body, &resquest)
	if err != nil {
		fmt.Println(err)
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	switch resquest.Function {
	case "QueryDeposit":
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
		findAllByCustID(requestGet, w)
	case "CreateDeposit":
		requestPost := models.RequestPost{}
		err = json.Unmarshal(body, &requestPost)
		if err != nil {
			fmt.Println(err)
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		err = requestPost.Validate()
		if err != nil {
			fmt.Println(err)
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		createDepositAward(requestPost, w)
	case "UpdateDeposit":
		requestUpdate := models.RequestUpdate{}
		err = json.Unmarshal(body, &requestUpdate)
		if err != nil {
			fmt.Println(err)
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		err = requestUpdate.Validate()
		if err != nil {
			fmt.Println(err)
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		updateDepositAward(requestUpdate, w)
	default:
		fmt.Println(err)
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("Cant read your function"))
		return
	}

}

func findAllByCustID(requestGet models.RequestGet, w http.ResponseWriter) {
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
		respGet.ResponseCode = http.StatusOK
		responses.JSON(w, http.StatusOK, &respGet)
		return
	}(thread)
}

func createDepositAward(requestPost models.RequestPost, w http.ResponseWriter) {
	db, err := database.Connect()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	thread := routine.NewAPIRoutine(db)
	func(postRepo repository.ApiRepository) {
		resPost, err := postRepo.CreateDeposit(&requestPost)
		if err != nil {
			fmt.Println(err)
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		resPost.Description = "Success"
		resPost.ResponseCode = http.StatusOK
		resPost.Data.CustomerId = requestPost.Data.CustomerID
		resPost.Data.LotteryCode = requestPost.Data.AccountDeposit
		responses.JSON(w, http.StatusOK, &resPost)
		return
	}(thread)
}

func updateDepositAward(requestUpdate models.RequestUpdate, w http.ResponseWriter) {

}
