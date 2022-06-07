package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	database, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgres")
	defer database.Close()

	//database, err := gorm.Open("postgres", dsn.String())

	if err != nil {
		panic("failed to connect to db")
	}
	database.AutoMigrate(&Book{})

	DB = database
}

// package models

// import (
// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/sqlite"
// )

// var DB *gorm.DB

// func ConnectDatabase() {
// 	database, err := gorm.Open("sqlite3", "test.db")

// 	if err != nil {
// 		panic("failed to connect to db")
// 	}
// 	database.AutoMigrate(&Book{})

// 	DB = database
// }
