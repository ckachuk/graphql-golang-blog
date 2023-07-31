package db

import (
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = ""
var DB *gorm.DB

func DBConnection() {
	var error error
	
	DB, error = gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})

	if error !=nil{
		log.Fatal(error)
	}else{
		log.Println("DB connected")
	}

}