package main 

import(
	// local
	"dairyapi/database"
	"dairyapi/model"

	// download
	"github.com/joho/godotenv"
    "log"
)



func main(){

	// load envs and db
	loadEnv()
	loadDatabase()
	
}

func loadEnv(){
	// load env files variables.
	err := godotenv.Load(".env")

	if err != nil{
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase(){
	// Connect to db and makemigrations.
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.User{})
}