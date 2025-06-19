package handler

import (
	"encoding/json"
	"github.com/baiv84/personio/model"
	"gorm.io/gorm"
	"net/http"
)

func create(pgConn *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var human model.Person

	err_json := json.NewDecoder(r.Body).Decode(&human)
	if err_json != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err_json.Error()))
		return
	}
	result := pgConn.Create(&human)
	if result.Error != nil {
		panic("Database connection error...")
	}
}
