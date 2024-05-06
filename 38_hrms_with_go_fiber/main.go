package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type MongoInstance struct{
	Client *mongo.Client
	Db	   *mongo.Database
}

var mg MongoInstance

const dbName = "hrms"

//const mongoURI = "mongodb+srv://ramgopalsiddh:Ram@8209820704%40Here@cluster0.6gq3p7r.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0" + dbName

type Employee struct{
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,idomitempty"`
	Name 	string `json:"name"`
	Salary 	float64 `json:"salary"`
	Age 	float64 `json:"age"`
}


func Connect() error {
	password := "password"
	encodedPassword := url.QueryEscape(password)

	mongoURI := "mongodb+srv://ramgopalsiddh:" + encodedPassword + "@cluster0.6gq3p7r.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

	fmt.Println(mongoURI)
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	db := client.Database(dbName)

	mg = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}



func main(){

	if err := Connect(); err !=nil {
		log.Fatal(err)
	}

	app := fiber.New()


	// Get all Employees
	app.Get("/employee", func(c *fiber.Ctx) error{

		query := bson.D{{}}

		cursor, err := mg.Db.Collection("employees").Find(c.Context(), query)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		var employees []Employee = make([]Employee, 0)

		if err := cursor.All(c.Context(), &employees); err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(employees)
	})


	// Creat a single Employee
	app.Post("/employee", func(c *fiber.Ctx) error {
		collection := mg.Db.Collection("employees")

		employee := new(Employee)

		if err := c.BodyParser(employee); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		employee.ID = primitive.NewObjectID()

		insertionResult, err := collection.InsertOne(c.Context(), employee)

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
		createdRecord := collection.FindOne(c.Context(), filter)

		createdEmployee := &Employee{}
		// decode record get from mongodb
		createdRecord.Decode(createdEmployee)

		return c.Status(201).JSON(createdEmployee)
	})


	// Update an employee
	app.Put("/employee/:id", func(c *fiber.Ctx) error {
		idParam := c.Params("id")

		employeeID, err := primitive.ObjectIDFromHex(idParam)

		if err != nil {
			return c.SendStatus(400)
		}

		employee := new(Employee)

		if err := c.BodyParser(employee); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		query := bson.D{{Key:"_id", Value: employeeID}}
		update := bson.D{
			{Key: "$set",
				Value: bson.D{
					{Key: "name", Value: employee.Name},
					{Key: "age", Value: employee.Age},
					{Key: "salary", Value: employee.Salary},
				},
			},
		}

		err = mg.Db.Collection("employees").FindOneAndUpdate(c.Context(), query, update).Err()
		if err != nil {
			if err == mongo.ErrNoDocuments{
				return c.SendStatus(400)
			}
			return c.SendStatus(500)
		}

		employee.ID = employeeID

		return c.Status(200).JSON(employee)
	})


	// Delete an employee
	app.Delete("/employee/:id", func(c *fiber.Ctx) error {
		employeeID, err := primitive.ObjectIDFromHex(c.Params("id"),)

		if err != nil {
			return c.SendStatus(400)
		}

		query := bson.D{{Key:"_id", Value: employeeID}}
		result, err := mg.Db.Collection("employees").DeleteOne(c.Context(), &query)

		if err != nil {
			return c.SendStatus(500)
		}

		if result.DeletedCount < 1{
			return c.SendStatus(404)
		}

		return c.Status(200).JSON("record deleted")
	})


	log.Fatal(app.Listen(":3000"))
}