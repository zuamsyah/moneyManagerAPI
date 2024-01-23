package handler

import (
	"encoding/json"
	"io/ioutil"
	"moneyManagerAPI/models"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func (idb *Idb) GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	startDate := r.FormValue("start")
	endDate := r.FormValue("end")

	transactions := []models.Transaction{}

	var err error
	if len(startDate) > 0 && len(endDate) > 0 {
		err = idb.DB.Model(&models.Transaction{}).Where("month = ?", time.Now().Month()).Where("created_at >= ? AND created_at <= ?", startDate, endDate).Find(&transactions).Error
	} else {
		err = idb.DB.Model(&models.Transaction{}).Where("month = ?", time.Now().Month()).Find(&transactions).Error
	}

	if err != nil {
		responseError := map[string]interface{}{"status": 500, "message": "Internal server error"}

		resultError, _ := json.Marshal(&responseError)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(resultError)
		return
	}

	category := models.Category{}
	for i, transaction := range transactions {
		idb.DB.First(&category, transaction.CategoryID)
		transactions[i].Category = category
	}

	response := models.ResponseTransactions{Status: 200, Data: transactions, Message: "Success Get Transactions"}

	result, _ := json.Marshal(&response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (idb *Idb) GetByIdTransactionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	transactionID := vars["id"]

	var transaction models.Transaction
	idb.DB.First(&transaction, transactionID)

	category := models.Category{}
	transaction.Category = category

	if transaction.ID == 0 {
		responseError := map[string]interface{}{"status": 404, "message": "Failed Get Transaction by ID"}

		resultError, _ := json.Marshal(&responseError)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resultError)
		return
	}

	response := map[string]interface{}{"status": 200, "data": transaction, "message": "Success Get Transaction by ID"}

	result, _ := json.Marshal(&response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (idb *Idb) CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	transaction := models.Transaction{}
	transaction.Month = int(time.Now().Month())
	reqBody, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(reqBody, &transaction)

	idb.DB.Create(&transaction)

	response := map[string]interface{}{"code": 201, "data": transaction, "message": "Success create data"}

	result, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (idb *Idb) UpdateByIdTransactionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	transactionID := vars["id"]

	transaction := models.Transaction{}
	reqBody, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(reqBody, &transaction)

	idb.DB.Model(&transaction).Where("id = ?", transactionID).Update(&transaction)
	idb.DB.First(&transaction, transactionID)

	response := map[string]interface{}{"status": 200, "data": transaction, "message": "Success update data"}

	result, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (idb *Idb) DeleteByIdTransactionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	transactionID := vars["id"]

	transaction := models.Transaction{}

	idb.DB.Delete(&transaction, transactionID)

	response := map[string]interface{}{"code": 200, "message": "Success delete data"}

	result, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(result)
}
