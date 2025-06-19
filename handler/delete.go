package handler

import (
	"github.com/baiv84/personio/model"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func deleteOne(pgConn *gorm.DB, w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic("Incorrect id")
	}
	human := model.Person{}
	pgConn.Unscoped().Delete(&human, id)
}
