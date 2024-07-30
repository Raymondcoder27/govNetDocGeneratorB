package initializers

import (
	"log"
	"os"

	// "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil{
		log.Fatal("Failed to connect to database.")
	}
}