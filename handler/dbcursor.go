package handler

import (
	"github.com/baiv84/personio/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

type DBCursor struct {
	pgConn *gorm.DB
}

func (cursor *DBCursor) InitDBEngine(dsn string) {
	var postgresDB *gorm.DB
	//const formatStr = "host=%s dbname=%s user=%s password=%s port=%s sslmode=disable"
	//dsn := fmt.Sprintf(formatStr, "localhost", "citizens", "postgres", "12345", "5432")
	postgresDB, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	err := postgresDB.AutoMigrate(&model.Person{})
	if err != nil {
		return
	}
	cursor.pgConn = postgresDB

}

func (cursor *DBCursor) Close() {
	dbInstance, _ := cursor.pgConn.DB()
	_ = dbInstance.Close()

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
