package database

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
    DBconn *gorm.DB
)

func Connect() {
    d, err := gorm.Open("postgres", "host=localhost port=5432 user=ram dbname=crm sslmode=disable")
    if err != nil {
        panic("Failed to connect to database: " + err.Error())
    }
    DBconn = d
}

func GetDB() *gorm.DB {
    if DBconn == nil {
        panic("Database connection is not initialized")
    }
    return DBconn
}
