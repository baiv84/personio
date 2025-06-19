package handler

import (
	"encoding/json"
	"github.com/baiv84/personio/model"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func update(pgConn *gorm.DB, w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic("Incorrect id")
	}

	var human model.Person
	err_json := json.NewDecoder(r.Body).Decode(&human)
	if err_json != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(err_json.Error()))
		if err != nil {
			return
		}
		return
	}

	elem := model.Person{}
	if len(human.FirstName) > 0 {
		pgConn.Model(&elem).Where("id = ?", id).Update("first_name", human.FirstName)
	}
	if len(human.SecondName) > 0 {
		pgConn.Model(&elem).Where("id = ?", id).Update("second_name", human.SecondName)
	}
	if len(human.ThirdName) > 0 {
		pgConn.Model(&elem).Where("id = ?", id).Update("third_name", human.ThirdName)
	}
	if len(human.Gender) > 0 {
		pgConn.Model(&elem).Where("id = ?", id).Update("gender", human.Gender)
	}
	if len(human.Country) > 0 {
		pgConn.Model(&elem).Where("id = ?", id).Update("country", human.Gender)
	}
	if human.Age > 0 {
		pgConn.Model(&elem).Where("id = ?", id).Update("age", human.Age)

	}

}
