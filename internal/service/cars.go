package service

import (
	"car-rental/internal/models"
	"car-rental/internal/repository"
	"context"
	"errors"
	"time"
)

type Carservice interface {
	GetCars(ctx context.Context) ([]models.Car, error)
	GetCarsByID(ctx context.Context, id uint64) (models.Car, error)
	CreateCar(ctx context.Context, car models.InputCar) (models.Car, error)
	EditCar(ctx context.Context, id uint64, car models.InputCar) (models.Car, error)
	DeleteCar(ctx context.Context, id uint64) (models.Car, error)
}
type carserviceImpl struct {
	carRepo repository.CarsQuery
}

func NewCarservice(carRepo repository.CarsQuery) Carservice {
	return &carserviceImpl{carRepo: carRepo}
}


func (s *carserviceImpl) GetCars(ctx context.Context) ([]models.Car, error) {
	cars, err := s.carRepo.GetCars(ctx)
	if err != nil {
		return nil, err
	}
	return cars, nil
}

func (s *carserviceImpl) GetCarsByID(ctx context.Context, id uint64) (models.Car, error) {
	car, err := s.carRepo.GetCarsByID(ctx, id)
	if err != nil {
		return models.Car{}, err
	}
	if car.ID == 0 {
		return models.Car{}, errors.New("car not found")
	}
	return car, nil
}

func (s *carserviceImpl) CreateCar(ctx context.Context, car models.InputCar) (models.Car, error) {
	NewCar := models.Car{}
	NewCar.Name = car.Name
	NewCar.Stock = car.Stock
	NewCar.DailyRent = car.DailyRent
	NewCar.CreatedAt = time.Now()

	// Call repoCarsitory to create car
	createdCar, err := s.carRepo.CreateCars(ctx, NewCar)
	if err != nil {
		return models.Car{}, err
	}
	return createdCar, nil
}

func (s *carserviceImpl) EditCar(ctx context.Context, id uint64, car models.InputCar) (models.Car, error) {
	updatedCar := models.Car{}
	updatedCar.Name = car.Name
	updatedCar.Stock = car.Stock
	updatedCar.DailyRent = car.DailyRent
	updatedCar.UpdatedAt = time.Now()

	// Call repoCarsitory to create car
	updatedCar, err := s.carRepo.EditCars(ctx, id, updatedCar)
	if err != nil {
		return models.Car{}, err
	}
	return updatedCar, nil
}

func (s *carserviceImpl) DeleteCar(ctx context.Context, id uint64) (models.Car, error) {
	car, err := s.carRepo.GetCarsByID(ctx, id)
	if err != nil {
		return models.Car{}, err
	}
	// if car doesn't exist, return
	if car.ID == 0 {
		return models.Car{}, nil
	}

	// delete car by id
	err = s.carRepo.DeleteCarsByID(ctx, id)
	if err != nil {
		return models.Car{}, err
	}

	return car, err
}
