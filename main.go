package main

import (
	"fmt"
	"log"
	"moneyManagerAPI/database"
	"moneyManagerAPI/handler"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	db := database.DBConnect()
	defer db.Close()

	idb := &handler.Idb{DB: db}

	r.HandleFunc("/auth_otp", handler.AuthOtpHandler).Methods("POST")

	r.HandleFunc("/categories", idb.GetCategoriesHandler).Methods("GET")
	r.HandleFunc("/categories/{id}", idb.GetByIdCategoryHandler).Methods("GET")
	r.HandleFunc("/categories", idb.CreateCategoryHandler).Methods("POST")
	r.HandleFunc("/categories/{id}", idb.UpdateByIdCategoryHandler).Methods("PUT")
	r.HandleFunc("/categories/{id}", idb.DeleteByIdCategoryHandler).Methods("DELETE")

	r.HandleFunc("/transactions", idb.GetTransactionsHandler).Methods("GET")
	r.HandleFunc("/transactions/{id}", idb.GetByIdTransactionHandler).Methods("GET")
	r.HandleFunc("/transactions", idb.CreateTransactionHandler).Methods("POST")
	r.HandleFunc("/transactions/{id}", idb.UpdateByIdTransactionHandler).Methods("PUT")
	r.HandleFunc("/transactions/{id}", idb.DeleteByIdTransactionHandler).Methods("DELETE")

	r.HandleFunc("/home", idb.GetHomeHandler).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:3000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("starting server :3000")
	log.Fatal(srv.ListenAndServe())
}
