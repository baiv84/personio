package handler

import (
	"encoding/json"
	"github.com/baiv84/personio/model"
	"gorm.io/gorm"
	"net/http"
)

func create(pgConn *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var human model.Person

	errJson := json.NewDecoder(r.Body).Decode(&human)
	if errJson != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(errJson.Error()))
		if err != nil {
			return
		}
		return
	}
	result := pgConn.Create(&human)
	if result.Error != nil {
		panic("Database connection error...")
	}

}
