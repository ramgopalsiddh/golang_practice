package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
) 

var (
	db *gorm.DB
)

func Connect() {
    d, err := gorm.Open("postgres", "host=localhost port=5432 user=ram dbname=bookmanagement sslmode=disable")
    if err != nil {
        panic("Failed to connect to database: " + err.Error())
    }
    db = d
}

func GetDB() *gorm.DB{
	return db
}