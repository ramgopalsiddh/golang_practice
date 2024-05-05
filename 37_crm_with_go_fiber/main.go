package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/ramgopalsiddh/crm_with_go-fiber/database"
	"github.com/ramgopalsiddh/crm_with_go-fiber/lead"
)


func setupRoutes(app *fiber.App){
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}


func initDatabase() {
	var err error
	database.DBconn, err = gorm.Open("postgres", "host=localhost port=5432 user=ram dbname=crm sslmode=disable")
	if err != nil {
		panic("Failed to connect database")
	}

	fmt.Println("Connected to database")
	database.DBconn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main(){
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBconn.Close()
}