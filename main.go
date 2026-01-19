package main

import (
	"go-fiber-postgre/dto"
	"go-fiber-postgre/internal/api"
	"go-fiber-postgre/internal/config"
	"go-fiber-postgre/internal/connection"
	"go-fiber-postgre/internal/repository"
	"go-fiber-postgre/internal/service"
	"net/http"

	jwtMid "github.com/gofiber/contrib/jwt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	dbConnection := connection.GetDatabase(cnf.Database)

	app := fiber.New()

	jwtMidd := jwtMid.New(jwtMid.Config{
		SigningKey : jwtMid.SigningKey{Key: []byte(cnf.Jwt.Key)},
		ErrorHandler: func (ctx *fiber.Ctx, err error) error {
			return ctx.Status(http.StatusUnauthorized).
				JSON(dto.CreateResponseError("Login terlebih dahulu"))			
		},
	})

	customerRepository := repository.NewCustomer(dbConnection)
	userRepository := repository.NewUser(dbConnection)

	customerService := service.NewCustomer(customerRepository)
	authService := service.NewAuth(cnf, userRepository)
	
	api.NewCustomer(app, customerService, jwtMidd)
	api.NewAuth(app, authService)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
