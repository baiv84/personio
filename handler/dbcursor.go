package handler

import (
	"gorm.io/gorm"
	"net/http"
)

type DBCursor struct {
	pgConn *gorm.DB
}

func (cursor *DBCursor) Init(pgConn *gorm.DB) {
	cursor.pgConn = pgConn
}

func (cursor *DBCursor) Create(w http.ResponseWriter, r *http.Request) {
	create(cursor.pgConn, w, r)
}

func (cursor *DBCursor) Read(w http.ResponseWriter, r *http.Request) {
	read(cursor.pgConn, w, r)
}

func (cursor *DBCursor) ReadOne(w http.ResponseWriter, r *http.Request) {
	readOne(cursor.pgConn, w, r)
}

func (cursor *DBCursor) Update(w http.ResponseWriter, r *http.Request) {
	update(cursor.pgConn, w, r)
}

func (cursor *DBCursor) Delete(w http.ResponseWriter, r *http.Request) {
	deleteOne(cursor.pgConn, w, r)
}
