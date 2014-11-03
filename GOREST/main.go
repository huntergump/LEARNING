package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// User Struct for Database returns //
type Users struct {
	Id       string
	Name     string
	Email    string
	Password string
}

// User Struct for JSON storage //
type Jusers struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Array or Slice of Users from Database //
type People []Users

func main() {

	// Setup the Database //
	db, err := gorm.Open("postgres", "user=HUNTER dbname=rhazes_users sslmode=disable")
	if err != nil {
		return
	}

	db.DB()

	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.SingularTable(true)

	// Prepare the Martini Server //
	server := martini.Classic()
	server.Get("/", func() string {
		return "hello world"
	})

	// Create a Martini Route //
	server.Get("/json/", func() string {

		// Pull all Database rows //
		var people People
		db.Find(&people)

		// Map Database returns into a JSON format //
		mappedjson := make(map[string]Jusers)

		for i := 0; i < len(people); i++ {
			mappedjson[fmt.Sprint(i)] = Jusers{Id: people[i].Id, Name: people[i].Name, Email: people[i].Email, Password: people[i].Password}
		}

		// Convert Mapped JSON data to a JSON String //
		jsonstring, err := json.Marshal(mappedjson)
		if err != nil {
			return string("JSON Marshal Error'd")
		}

		// Return JSON Data //
		return string(jsonstring)
	})

	server.Use(martini.Static("app"))
	server.Run()
}
