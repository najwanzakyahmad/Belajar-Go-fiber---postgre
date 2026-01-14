package main

import (
	"go-fiber-postgre/internal/config"
	"go-fiber-postgre/internal/connection"
	"go-fiber-postgre/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New()

	customerRepository := repository.NewCustomer(dbConnection)
	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)

	
}
