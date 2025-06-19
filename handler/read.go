package handler

import (
	"encoding/json"
	"github.com/baiv84/personio/model"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func read(pgConn *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var humans []model.Person
	pgConn.Find(&humans)
	resp, _ := json.Marshal(humans)
	w.Write(resp)

}

func readOne(pgConn *gorm.DB, w http.ResponseWriter, r *http.Request) {
	human := model.Person{}
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)
	pgConn.Select("id", "first_name", "second_name", "third_name", "country", "gender", "age").
		Where("id=?", id).Find(&human)

	resp, _ := json.Marshal(human)
	w.Write(resp)

}
