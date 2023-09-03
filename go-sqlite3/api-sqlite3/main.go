package main

import (
	"log"
	"net/http"
	"personweb/models"

	"github.com/gin-gonic/gin"
)

// re-usable function to check errors
func checkErr(err error){
	if err != nil{
		log.Fatal(err)
	}
}

func main(){
	r := gin.Default()
	err := models.ConnectDb()
	checkErr(err)
	// API v1
	v1 := r.Group("/api/v1")
	{
	  v1.GET("person", getPersons)
	  v1.GET("person/:id", getPersonById)
	  v1.POST("person", addPerson)
	  v1.PUT("person/:id", updatePerson)
	  v1.DELETE("person/:id", deletePerson)
	  v1.OPTIONS("person", options)
	}


	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	r.Run()
}


// Functions for routes.
func getPersons(c *gin.Context){
	persons,err := models.GetPersons(10)

	checkErr(err)

	if persons == nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	}else{
		// return persons
		c.JSON(http.StatusOK, gin.H{"data": persons})
	}
	
}

func getPersonById(c *gin.Context){
	// take id parameter
	id := c.Param("id")

	person, err := models.GetPersonById(id)

	checkErr(err)

	// if firstname is blank we can assume nothing was found.
	if person.FirstName == ""{
		c.JSON(http.StatusBadRequest,gin.H{"error":"No Records Found"})
		return
	}else{
		c.JSON(http.StatusOK,gin.H{"data":person})
	}

	// c.JSON(http.StatusOK,gin.H{"message":"getPersonById" + id + " called"})
}

func addPerson(c *gin.Context){

	// return json test response
	c.JSON(http.StatusOK,gin.H{"message":"addPerson Called"})
}

func updatePerson(c *gin.Context){

	// return json test response
	c.JSON(http.StatusOK,gin.H{"message":"updatePerson Called"})
}


func deletePerson(c *gin.Context){
	id := c.Param("id")

	// return json test response
	c.JSON(http.StatusOK,gin.H{"message":"deletePerson of id" + id + " called"})
}

func options(c *gin.Context){
	// return json test response
	c.JSON(http.StatusOK,gin.H{"message":"options Called"})
}