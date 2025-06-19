package model

import "gorm.io/gorm"

type PersonPure struct {
	FirstName  string `gorm:"type:varchar(255)"`
	SecondName string `gorm:"type:varchar(255)"`
	ThirdName  string `gorm:"type:varchar(255)"`
	Gender     string `gorm:"type:varchar(7)"`
	Age        int    `gorm:"type:int4"`
	Country    string `gorm:"type:varchar(5)"`
}

type Person struct {
	gorm.Model
	PersonPure
}
