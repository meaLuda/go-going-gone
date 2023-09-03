package models

import (
	"database/sql"
	"strconv"
	_ "github.com/mattn/go-sqlite3"
)

// global variable
var DB *sql.DB

func ConnectDb() error{
	db, err := sql.Open("sqlite3","models/names.db")

	// check when an error occurs
	if err != nil{
		return err
	}

	// Set value of DB to this new db
	DB = db 
	return nil
}

// Person model
type Person struct{
	Id int `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email string   `json:"email"`
	IpAddress string `json:"ip_address"`
}



// Queries
func GetPersons(count int) ([]Person, error) {

	rows, err := DB.Query("SELECT * FROM people LIMIT " + strconv.Itoa(count))
	// rows, err := DB.Query("SELECT id, first_name, last_name, email, ip_address from people LIMIT " + strconv.Itoa(count))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	people := make([]Person, 0)

	for rows.Next() {
		singlePerson := Person{}
		err = rows.Scan(&singlePerson.Id, &singlePerson.FirstName, &singlePerson.LastName, &singlePerson.Email, &singlePerson.IpAddress)

		if err != nil {
			return nil, err
		}

		people = append(people, singlePerson)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return people, err
}


// Get Person by id
func GetPersonById(id string) (Person, error){
	stmt, err := DB.Prepare("SELECT * FROM people WHERE id = ?")

	if err != nil {
		return Person{}, err
	}

	person := Person{}

	newErr := stmt.QueryRow(id).Scan(&person.Id,&person.FirstName,&person.LastName,&person.Email, &person.IpAddress)

	if newErr != nil{
		if newErr == sql.ErrNoRows{
			return Person{},err
		}
		return Person{},newErr
	}

	return person,nil

}