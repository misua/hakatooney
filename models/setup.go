package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// "log"
	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	database, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgres")

	if err != nil {
		panic("failed to connect to db")

	}
	defer database.Close()
	database.AutoMigrate(&Book{})

	//return database

	DB = database
}

// docker run --name basic-postgres --rm -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=postgres -e PGDATA=/var/lib/postgresql/data/pgdata -v /tmp:/var/lib/postgresql/data -p 5432:5432 -it postgres:14.1-alpine
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
