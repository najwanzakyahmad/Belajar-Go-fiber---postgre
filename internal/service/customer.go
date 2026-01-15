package service

import (
	"context"
	"database/sql"
	"go-fiber-postgre/domain"
	"go-fiber-postgre/dto"
	"time"

	"github.com/google/uuid"
)

type customerService struct {
	customerRepository domain.CustomerRepository
}

// Index implements [domain.CustomerService].
func (c *customerService) Index(ctx context.Context) ([]dto.CustomerData, error) {
	customers, err := c.customerRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var customerData []dto.CustomerData
	for _, v := range customers {
		customerData = append(customerData, dto.CustomerData{
			ID:   v.ID,
			Code: v.Code,
			Name: v.Name,
		})
	}

	return customerData, nil
}

// Create implements [domain.CustomerService].
func (c *customerService) Create(ctx context.Context, req dto.CreateCustomerRequest) error {
	customer := domain.Customer{
		ID: uuid.NewString(),
		Code: req.Code,
		Name: req.Name,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}

	return c.customerRepository.Save(ctx, &customer)
}


func NewCustomer(customerRepository domain.CustomerRepository) domain.CustomerService {
	return &customerService{
		customerRepository: customerRepository,
	}
}
