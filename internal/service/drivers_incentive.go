package service

import (
	"car-rental/internal/models"
	"car-rental/internal/repository"
	"context"
	"errors"
	"time"
)

type DriversIncentiveservice interface {
	GetDriversIncentive(ctx context.Context) ([]models.DriverIncentive, error)
	GetDriversIncentiveByID(ctx context.Context, id uint64) (models.DriverIncentive, error)
	CreateDriverIncentive(ctx context.Context, driverIncentive models.InputDriverIncentive) (models.DriverIncentive, error)
	EditDriverIncentive(ctx context.Context, id uint64, driverIncentive models.InputDriverIncentive) (models.DriverIncentive, error)
	DeleteDriverIncentive(ctx context.Context, id uint64) (models.DriverIncentive, error)
}
type driversIncentiveerviceImpl struct {
	driverIncentiveRepo repository.DriversIncentiveQuery
}

func NewDriversIncentiveervice(driverIncentiveRepo repository.DriversIncentiveQuery) DriversIncentiveservice {
	return &driversIncentiveerviceImpl{driverIncentiveRepo: driverIncentiveRepo}
}


func (s *driversIncentiveerviceImpl) GetDriversIncentive(ctx context.Context) ([]models.DriverIncentive, error) {
	driversIncentive, err := s.driverIncentiveRepo.GetDriversIncentive(ctx)
	if err != nil {
		return nil, err
	}
	return driversIncentive, nil
}

func (s *driversIncentiveerviceImpl) GetDriversIncentiveByID(ctx context.Context, id uint64) (models.DriverIncentive, error) {
	driverIncentive, err := s.driverIncentiveRepo.GetDriversIncentiveByID(ctx, id)
	if err != nil {
		return models.DriverIncentive{}, err
	}
	if driverIncentive.ID == 0 {
		return models.DriverIncentive{}, errors.New("driverIncentive not found")
	}
	return driverIncentive, nil
}

func (s *driversIncentiveerviceImpl) CreateDriverIncentive(ctx context.Context, driverIncentive models.InputDriverIncentive) (models.DriverIncentive, error) {
	NewDriverIncentive := models.DriverIncentive{}
	NewDriverIncentive.BookingID = &driverIncentive.BookingID
	NewDriverIncentive.Incentive = driverIncentive.Incentive
	NewDriverIncentive.CreatedAt = time.Now()

	createdDriverIncentive, err := s.driverIncentiveRepo.CreateDriversIncentive(ctx, NewDriverIncentive)
	if err != nil {
		return models.DriverIncentive{}, err
	}
	return createdDriverIncentive, nil
}

func (s *driversIncentiveerviceImpl) EditDriverIncentive(ctx context.Context, id uint64, driverIncentive models.InputDriverIncentive) (models.DriverIncentive, error) {
	updatedDriverIncentive := models.DriverIncentive{}
	updatedDriverIncentive.BookingID = &driverIncentive.BookingID
	updatedDriverIncentive.Incentive = driverIncentive.Incentive
	updatedDriverIncentive.UpdatedAt = time.Now()

	updatedDriverIncentive, err := s.driverIncentiveRepo.EditDriversIncentive(ctx, id, updatedDriverIncentive)
	if err != nil {
		return models.DriverIncentive{}, err
	}
	return updatedDriverIncentive, nil
}

func (s *driversIncentiveerviceImpl) DeleteDriverIncentive(ctx context.Context, id uint64) (models.DriverIncentive, error) {
	driverIncentive, err := s.driverIncentiveRepo.GetDriversIncentiveByID(ctx, id)
	if err != nil {
		return models.DriverIncentive{}, err
	}
	if driverIncentive.ID == 0 {
		return models.DriverIncentive{}, nil
	}

	err = s.driverIncentiveRepo.DeleteDriversIncentiveByID(ctx, id)
	if err != nil {
		return models.DriverIncentive{}, err
	}

	return driverIncentive, err
}
