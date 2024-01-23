package handler

import (
	"encoding/json"
	"io/ioutil"
	"moneyManagerAPI/models"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Idb struct {
	DB *gorm.DB
}

func (idb *Idb) GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	categories := []models.Category{}

	idb.DB.Find(&categories)

	response := map[string]interface{}{"code": 200, "data": categories, "message": "Success Get Categories"}

	results, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func (idb *Idb) GetByIdCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID := vars["id"]

	var category models.Category
	idb.DB.First(&category, categoryID)

	response := map[string]interface{}{"code": 200, "data": category, "message": "Success Get Category by ID"}

	result, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (idb *Idb) CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
	category := models.Category{}
	reqBody, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(reqBody, &category)

	idb.DB.Create(&category)

	response := map[string]interface{}{"code": 201, "data": category, "message": "Success create data"}

	result, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (idb *Idb) UpdateByIdCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID := vars["id"]

	category := models.Category{}
	reqBody, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(reqBody, &category)

	idb.DB.Model(&category).Where("id = ?", categoryID).Update(&category)
	idb.DB.First(&category, categoryID)

	response := map[string]interface{}{"status": 200, "data": category, "message": "Success update data"}

	result, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func (idb *Idb) DeleteByIdCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryID := vars["id"]

	category := models.Category{}

	idb.DB.Delete(&category, categoryID)

	response := map[string]interface{}{"code": 200, "message": "Success delete data"}

	result, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(result)
}
