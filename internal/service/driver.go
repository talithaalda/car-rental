package service

import (
	"car-rental/internal/models"
	"car-rental/internal/repository"
	"context"
	"errors"
	"time"
)

type Driverservice interface {
	GetDrivers(ctx context.Context) ([]models.Driver, error)
	GetDriversByID(ctx context.Context, id uint64) (models.Driver, error)
	CreateDriver(ctx context.Context, driver models.InputDriver) (models.Driver, error)
	EditDriver(ctx context.Context, id uint64, driver models.InputDriver) (models.Driver, error)
	DeleteDriver(ctx context.Context, id uint64) (models.Driver, error)
}
type driverserviceImpl struct {
	driverRepo repository.DriversQuery
}

func NewDriverservice(driverRepo repository.DriversQuery) Driverservice {
	return &driverserviceImpl{driverRepo: driverRepo}
}


func (s *driverserviceImpl) GetDrivers(ctx context.Context) ([]models.Driver, error) {
	drivers, err := s.driverRepo.GetDrivers(ctx)
	if err != nil {
		return nil, err
	}
	return drivers, nil
}

func (s *driverserviceImpl) GetDriversByID(ctx context.Context, id uint64) (models.Driver, error) {
	driver, err := s.driverRepo.GetDriversByID(ctx, id)
	if err != nil {
		return models.Driver{}, err
	}
	if driver.ID == 0 {
		return models.Driver{}, errors.New("driver not found")
	}
	return driver, nil
}

func (s *driverserviceImpl) CreateDriver(ctx context.Context, driver models.InputDriver) (models.Driver, error) {
	NewDriver := models.Driver{}
	NewDriver.Name = driver.Name
	NewDriver.NIK = driver.NIK
	NewDriver.Phone = driver.Phone
	NewDriver.DailyCost = driver.DailyCost
	NewDriver.CreatedAt = time.Now()

	// Call repoDriversitory to create driver
	createdDriver, err := s.driverRepo.CreateDrivers(ctx, NewDriver)
	if err != nil {
		return models.Driver{}, err
	}
	return createdDriver, nil
}

func (s *driverserviceImpl) EditDriver(ctx context.Context, id uint64, driver models.InputDriver) (models.Driver, error) {
	updatedDriver := models.Driver{}
	updatedDriver.Name = driver.Name
	updatedDriver.NIK = driver.NIK
	updatedDriver.Phone = driver.Phone
	updatedDriver.DailyCost = driver.DailyCost
	updatedDriver.UpdatedAt = time.Now()

	// Call repoDriversitory to create driver
	updatedDriver, err := s.driverRepo.EditDrivers(ctx, id, updatedDriver)
	if err != nil {
		return models.Driver{}, err
	}
	return updatedDriver, nil
}

func (s *driverserviceImpl) DeleteDriver(ctx context.Context, id uint64) (models.Driver, error) {
	driver, err := s.driverRepo.GetDriversByID(ctx, id)
	if err != nil {
		return models.Driver{}, err
	}
	// if driver doesn't exist, return
	if driver.ID == 0 {
		return models.Driver{}, nil
	}

	// delete driver by id
	err = s.driverRepo.DeleteDriversByID(ctx, id)
	if err != nil {
		return models.Driver{}, err
	}

	return driver, err
}
