package service

import (
	"car-rental/internal/models"
	"car-rental/internal/repository"
	"context"
	"errors"
	"time"
)

type CustomerService interface {
	GetCustomers(ctx context.Context) ([]models.Customer, error)
	GetCustomersByID(ctx context.Context, id uint64) (models.Customer, error)
	CreateCustomer(ctx context.Context, customer models.InputCustomer) (models.Customer, error)
	EditCustomer(ctx context.Context, id uint64, customer models.InputCustomer) (models.Customer, error)
	DeleteCustomer(ctx context.Context, id uint64) (models.Customer, error)
}
type customerServiceImpl struct {
	customerRepo repository.CustomersQuery
}

func NewCustomerService(customerRepo repository.CustomersQuery) CustomerService {
	return &customerServiceImpl{customerRepo: customerRepo}
}


func (s *customerServiceImpl) GetCustomers(ctx context.Context) ([]models.Customer, error) {
	customers, err := s.customerRepo.GetCustomers(ctx)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (s *customerServiceImpl) GetCustomersByID(ctx context.Context, id uint64) (models.Customer, error) {
	customer, err := s.customerRepo.GetCustomersByID(ctx, id)
	if err != nil {
		return models.Customer{}, err
	}
	if customer.ID == 0 {
		return models.Customer{}, errors.New("customer not found")
	}
	return customer, nil
}

func (s *customerServiceImpl) CreateCustomer(ctx context.Context, customer models.InputCustomer) (models.Customer, error) {
	NewCustomer := models.Customer{}
	NewCustomer.Name = customer.Name
	NewCustomer.NIK = customer.NIK
	NewCustomer.Phone = customer.Phone
	NewCustomer.CreatedAt = time.Now()

	// Call repoCustomersitory to create customer
	createdCustomer, err := s.customerRepo.CreateCustomers(ctx, NewCustomer)
	if err != nil {
		return models.Customer{}, err
	}
	return createdCustomer, nil
}

func (s *customerServiceImpl) EditCustomer(ctx context.Context, id uint64, customer models.InputCustomer) (models.Customer, error) {
	updatedCustomer := models.Customer{}
	updatedCustomer.Name = customer.Name
	updatedCustomer.NIK = customer.NIK
	updatedCustomer.Phone = customer.Phone
	updatedCustomer.UpdatedAt = time.Now()

	// Call repoCustomersitory to create customer
	updatedCustomer, err := s.customerRepo.EditCustomers(ctx, id, updatedCustomer)
	if err != nil {
		return models.Customer{}, err
	}
	return updatedCustomer, nil
}

func (s *customerServiceImpl) DeleteCustomer(ctx context.Context, id uint64) (models.Customer, error) {
	customer, err := s.customerRepo.GetCustomersByID(ctx, id)
	if err != nil {
		return models.Customer{}, err
	}
	// if customer doesn't exist, return
	if customer.ID == 0 {
		return models.Customer{}, nil
	}

	// delete customer by id
	err = s.customerRepo.DeleteCustomersByID(ctx, id)
	if err != nil {
		return models.Customer{}, err
	}

	return customer, err
}
