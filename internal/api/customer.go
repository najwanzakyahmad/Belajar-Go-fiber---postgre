package api

import (
	"context"
	"go-fiber-postgre/domain"
	"go-fiber-postgre/dto"
	"go-fiber-postgre/internal/util"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type customerApi struct {
	customerService domain.CustomerService
}


func NewCustomer(app *fiber.App, customerService domain.CustomerService) {
	ca := customerApi {
		customerService: customerService,
	}

	app.Get("/customers", ca.Index)
	app.Get("/customer/:id", ca.Show)
	app.Post("/customers", ca.Create)
	app.Put("/customer/:id", ca.Update)
	app.Delete("/customer/:id", ca.Delete)
}

func (ca customerApi) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()
	
	res, err := ca.customerService.Index(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.JSON(dto.CreateResponseSuccess(res))
}

func (ca customerApi) Show(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")
	result, err := ca.customerService.Show(c, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.JSON(dto.CreateResponseSuccess(result))
}
 
func (ca customerApi) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.CreateCustomerRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadGateway).
			JSON(dto.CreateResponseErrorData("validation failed", fails))
	}

	err := ca.customerService.Create(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.Status(http.StatusCreated).
		JSON(dto.CreateResponseSuccess(""))
}

func (ca customerApi) Update(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.UpdateCustomerRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := util.Validate(req)
	if len(fails) > 0 {
		return ctx.Status(http.StatusBadGateway).
			JSON(dto.CreateResponseErrorData("validation failed", fails))
	}

	req.ID = ctx.Params("id")
	err := ca.customerService.Update(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	} 

	return ctx.Status(http.StatusOK).
		JSON(dto.CreateResponseSuccess(""))
}

func (ca customerApi) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	id := ctx.Params("id")
	err := ca.customerService.Delete(c, id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}

	return ctx.SendStatus(http.StatusNoContent)
}