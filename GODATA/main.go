package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type Users struct {
	Id       int64
	Name     string
	Email    string
	Password string
}

type People []Users

func main() {
	db, err := gorm.Open("postgres", "user=HUNTER dbname=rhazes_users sslmode=disable")
	if err != nil {
		return
	}

	db.DB()

	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	db.SingularTable(true)

	// Print "people"returns all database entries //
	// Print "people[1]" returns on the one entry //
	// Print "people[2]" returns the Name within the one entry //
	var people People
	db.Find(&people)
	fmt.Println(people[1].Name)

}
