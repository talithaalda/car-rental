package repository

import (
	"context"

	"car-rental/internal/infrastructure"
	"car-rental/internal/models"

	"gorm.io/gorm"
)

type CustomersQuery interface {
	GetCustomers(ctx context.Context) ([]models.Customer, error)
	GetCustomersByID(ctx context.Context, id uint64) (models.Customer, error)
	EditCustomers(ctx context.Context, id uint64, customers models.Customer) (models.Customer, error)
	DeleteCustomersByID(ctx context.Context, id uint64) error
	CreateCustomers(ctx context.Context, customers models.Customer) (models.Customer, error)
	DeleteMembershipByCustomer(ctx context.Context, id uint64, customer models.Customer) (models.Customer, error)
}

type CustomersCommand interface {
	CreateCustomers(ctx context.Context, customers models.Customer) (models.Customer, error)
}

type customersQueryImpl struct {
	db infrastructure.GormPostgres
}

func NewCustomersQuery(db infrastructure.GormPostgres) CustomersQuery {
	return &customersQueryImpl{db: db}
}

func (u *customersQueryImpl) GetCustomers(ctx context.Context) ([]models.Customer, error) {
	db := u.db.GetConnection()
	customers := []models.Customer{}

	if err := db.WithContext(ctx).
		Preload("Membership").     
		Find(&customers).Error; err != nil {
		return nil, err
	}

	return customers, nil
}

func (u *customersQueryImpl) GetCustomersByID(ctx context.Context, id uint64) (models.Customer, error) {
	db := u.db.GetConnection()
	customers := models.Customer{}

	if err := db.WithContext(ctx).
		Preload("Membership").    
		First(&customers, id).Error; err != nil { 
		if err == gorm.ErrRecordNotFound {
			return models.Customer{}, nil
		}
		return models.Customer{}, err
	}

	return customers, nil
}

func (u *customersQueryImpl) DeleteCustomersByID(ctx context.Context, id uint64) error {
	db := u.db.GetConnection()
	customer := models.Customer{}
	if err := db.WithContext(ctx).
		Preload("Membership").
		First(&customer, id).Error; err != nil {
		return err
	}

	if err := db.WithContext(ctx).
		Delete(&customer).Error; err != nil {
		return err
	}

	return nil
}

func (u *customersQueryImpl) CreateCustomers(ctx context.Context, customers models.Customer) (models.Customer, error) {
	
	db := u.db.GetConnection()
	if err := db.
		WithContext(ctx).
		Table("customers").
		Save(&customers).Error; err != nil {
		return models.Customer{}, err
	}
	return customers, nil
}
func (u *customersQueryImpl) EditCustomers(ctx context.Context, id uint64, customer models.Customer) (models.Customer, error) {
	db := u.db.GetConnection()

	if err := db.WithContext(ctx).
		Model(&models.Customer{}).
		Where("id = ?", id).
		Updates(customer).Error; err != nil {
		return models.Customer{}, err
	}


	updatedCustomers := models.Customer{}
	if err := db.
		WithContext(ctx).
		Preload("Membership").
		First(&updatedCustomers, id).Error; err != nil {
		return models.Customer{}, err
	}
	return updatedCustomers, nil
}

func (u *customersQueryImpl) DeleteMembershipByCustomer(ctx context.Context, id uint64, customer models.Customer) (models.Customer, error) {
	db := u.db.GetConnection()

	if err := db.Model(&models.Customer{}).
		Where("id = ?", id).
		Update("membership_id", nil).Error; err != nil {
		return models.Customer{}, err
	}

	return customer, nil
}
