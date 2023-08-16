package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	name string
	age  int
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("Can't migrate User scheme")
	}
	db.Create(&User{name: "Moein", age: 20})
}
