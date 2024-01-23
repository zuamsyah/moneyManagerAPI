package handler

import (
	"encoding/json"
	"moneyManagerAPI/models"
	"net/http"
)

func (idb *Idb) GetHomeHandler(w http.ResponseWriter, r *http.Request) {
	home := models.Home{}

	idb.DB.Find(&home)

	response := map[string]interface{}{"status": 200, "data": home, "message": "Success Get home"}

	result, _ := json.Marshal(&response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
