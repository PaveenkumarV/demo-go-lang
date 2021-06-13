package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "user:password@tcp/first_go?charset=utf8mb4,utf8"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	users := user{name: "Harish", email: "goutham@gmail", phone: 790789, address: "Bangalore"}
	db.Create(users) // Creating

	db.Where("name = ?", "Goutham") // Reading

	users.name = "Goutham" // Updating
	users.phone = 100
	db.Save(&users)

	db.Where("name = ?", "Harish").Delete(&users) // Deleting

}

type user struct {
	name    string
	email   string
	phone   int
	address string
}
