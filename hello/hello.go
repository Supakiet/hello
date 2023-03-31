package main

import (
	"fmt"
	"log"

	"github.com/supakiet/greetings"
)

// var db *gorm.DB

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)
}

// type Book struct {
// 	gorm.Model
// 	Name   string `json:"name"`
// 	Author string `json:"author"`
// }

// func NewBook(c *gin.Context) {
// 	var book Book
// 	if err := c.Bind(&book); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 	}

// 	result := db.Create(&book)
// 	if err := result.Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 	}

// 	c.Status(http.StatusCreated)
// }
