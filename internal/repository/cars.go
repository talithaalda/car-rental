package repository

import (
	"context"

	"car-rental/internal/infrastructure"
	"car-rental/internal/models"

	"gorm.io/gorm"
)

type CarsQuery interface {
	GetCars(ctx context.Context) ([]models.Car, error)
	GetCarsByID(ctx context.Context, id uint64) (models.Car, error)
	EditCars(ctx context.Context, id uint64, cars models.Car) (models.Car, error)
	DeleteCarsByID(ctx context.Context, id uint64) error
	CreateCars(ctx context.Context, cars models.Car) (models.Car, error)
}

type CarsCommand interface {
	CreateCars(ctx context.Context, cars models.Car) (models.Car, error)
}

type carsQueryImpl struct {
	db infrastructure.GormPostgres
}

func NewCarsQuery(db infrastructure.GormPostgres) CarsQuery {
	return &carsQueryImpl{db: db}
}

func (u *carsQueryImpl) GetCars(ctx context.Context) ([]models.Car, error) {
	db := u.db.GetConnection()
	cars := []models.Car{}
	if err := db.
		WithContext(ctx).
		Table("cars").
		Find(&cars).Error; err != nil {
		return nil, err
	}
	return cars, nil
}

func (u *carsQueryImpl) GetCarsByID(ctx context.Context, id uint64) (models.Car, error) {
	db := u.db.GetConnection()
	cars := models.Car{}
	if err := db.
		WithContext(ctx).
		Table("cars").
		Where("id = ?", id).
		Find(&cars).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Car{}, nil
		}

		return models.Car{}, err
	}
	return cars, nil
}

func (u *carsQueryImpl) DeleteCarsByID(ctx context.Context, id uint64) error {
	db := u.db.GetConnection()
	if err := db.
		WithContext(ctx).
		Table("cars").
		Delete(&models.Car{ID: uint(id)}).
		Error; err != nil {
		return err
	}
	return nil
}

func (u *carsQueryImpl) CreateCars(ctx context.Context, cars models.Car) (models.Car, error) {
	
	db := u.db.GetConnection()
	if err := db.
		WithContext(ctx).
		Table("cars").
		Save(&cars).Error; err != nil {
		return models.Car{}, err
	}
	return cars, nil
}
func (u *carsQueryImpl) EditCars(ctx context.Context, id uint64, car models.Car) (models.Car, error) {
	db := u.db.GetConnection()
	updatedCars := models.Car{}
	if err := db.
		WithContext(ctx).
		Table("cars").
		Where("id = ?", id).Updates(&car).First(&updatedCars).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return models.Car{}, nil
			}
		}
	return updatedCars, nil
}
