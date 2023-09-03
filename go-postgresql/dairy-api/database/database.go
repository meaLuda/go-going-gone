package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)


// Create universal db
var Database *gorm.DB 


func Connect(){
	// Connect to db
	var err error 

	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	databasename := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// Uniform Resource Identifier
	uri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", host, username, password, databasename, port)

	Database, err = gorm.Open(postgres.Open(uri),&gorm.Config{})

	if err != nil{
		panic(err)
	}else{
		fmt.Println("Successfully connected to the database")
	}
}
