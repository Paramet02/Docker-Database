package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

const (
	host         = "localhost"
	port         = 5432
	databaseName = "mydatabase"
	username     = "myuser"
	password     = "mypassword"
)

var db *sql.DB

type Products struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
}

func main() {
	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, databaseName)

	sdb, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err) 
	}
	
	db = sdb
	
	// defer is a command that runs before the last command in a program (or function). 
	// It is usually used for cleanup commands to completely shut down a process before it stops working.
	defer db.Close()

	err = db.Ping()

	if err != nil {
		log.Fatal(err) 
	}
	
	// connection database successful
	fmt.Println("connection database successful")

	app := fiber.New()
	
	app.Get("/product/:id", getHandler)
	app.Post("/product", postHandler)
	app.Put("/product/:id", updateHandler)
	app.Delete("/product/:id", delHandler)
	app.Get("/products", getsHandler)

	app.Listen(":8080")
	
	// ----------------------------------------------------------------//
	// err = createDatabase(&Products{Name: "Go Product 3 ", Price: 45})
	// if err != nil {
	// 	log.Fatal(err) 
	// }

	// Create successful
	// fmt.Println("Create database successful")
	// ----------------------------------------------------------------//

	// ---------------------------------------------------------------//
	// p ,err := getProducts(2)
	// if err != nil {
	// 	log.Fatal(err) 
	// }

	// fmt.Println("Get database successful!", p)
	// ----------------------------------------------------------------//

	// ----------------------------------------------------------------//
	// err = updateProducts(3, &Products{Name: "XYZ", Price: 200})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	
	// fmt.Println("update database successful!")
	// ----------------------------------------------------------------//

	// ----------------------------------------------------------------//
	// err = delProducts(2)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	
	// fmt.Println("del database successful!")
	// ----------------------------------------------------------------//

}

func getHandler(c *fiber.Ctx) error {
	
	productsId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	products, err := getProduct(productsId)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(products)
}

func postHandler(c *fiber.Ctx) error {
	// Reserve space in adders first.
	p := new(Products)
	if err := c.BodyParser(p) ; err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	// Insert product into database 
	err := createDatabase(p)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(p)	
	
}

func updateHandler(c *fiber.Ctx) error {
	// change string to int by strconv
	updateId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Reserve space in adders first.
	p := new(Products)
	if err := c.BodyParser(p); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// update products into database
	if err := updateProducts(updateId, p); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	
	return c.JSON(p)
}

func delHandler(c *fiber.Ctx) error {
	delId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := delProducts(delId); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	
	return c.SendStatus(fiber.StatusNoContent)
}

func getsHandler(c *fiber.Ctx) error {
	gets, err := getProducts()
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(gets)
}