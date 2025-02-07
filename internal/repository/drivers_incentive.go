package repository

import (
	"context"

	"car-rental/internal/infrastructure"
	"car-rental/internal/models"

	"gorm.io/gorm"
)

type DriversIncentiveQuery interface {
	GetDriversIncentive(ctx context.Context) ([]models.DriverIncentive, error)
	GetDriversIncentiveByID(ctx context.Context, id uint64) (models.DriverIncentive, error)
	EditDriversIncentive(ctx context.Context, id uint64, driversIncentive models.DriverIncentive) (models.DriverIncentive, error)
	DeleteDriversIncentiveByID(ctx context.Context, id uint64) error
	CreateDriversIncentive(ctx context.Context, driversIncentive models.DriverIncentive) (models.DriverIncentive, error)
}

type DriversIncentiveCommand interface {
	CreateDriversIncentive(ctx context.Context, driversIncentive models.DriverIncentive) (models.DriverIncentive, error)
}

type driversIncentiveQueryImpl struct {
	db infrastructure.GormPostgres
}

func NewDriversIncentiveQuery(db infrastructure.GormPostgres) DriversIncentiveQuery {
	return &driversIncentiveQueryImpl{db: db}
}

func (u *driversIncentiveQueryImpl) GetDriversIncentive(ctx context.Context) ([]models.DriverIncentive, error) {
	db := u.db.GetConnection()
	driversIncentive := []models.DriverIncentive{}

	if err := db.WithContext(ctx).
		Preload("Booking").
		Preload("Booking.Customer.Membership").
		Preload("Booking.Customer").
		Preload("Booking.Car").
		Preload("Booking.Driver").
		Preload("Booking.BookingType").
		Find(&driversIncentive).Error; err != nil {
		return nil, err
	}

	return driversIncentive, nil
}

func (u *driversIncentiveQueryImpl) GetDriversIncentiveByID(ctx context.Context, id uint64) (models.DriverIncentive, error) {
	db := u.db.GetConnection()
	driversIncentive := models.DriverIncentive{}

	if err := db.WithContext(ctx).
		Preload("Booking").
		Preload("Booking.Customer.Membership").
		Preload("Booking.Customer").
		Preload("Booking.Car").
		Preload("Booking.Driver").
		Preload("Booking.BookingType").
		First(&driversIncentive, id).Error; err != nil { 
		if err == gorm.ErrRecordNotFound {
			return models.DriverIncentive{}, nil
		}
		return models.DriverIncentive{}, err
	}

	return driversIncentive, nil
}

func (u *driversIncentiveQueryImpl) DeleteDriversIncentiveByID(ctx context.Context, id uint64) error {
	db := u.db.GetConnection()
	if err := db.
		WithContext(ctx).
		Table("driver_incentives").
		Delete(&models.DriverIncentive{ID: uint(id)}).
		Error; err != nil {
		return err
	}
	return nil
}

func (u *driversIncentiveQueryImpl) CreateDriversIncentive(ctx context.Context, driversIncentive models.DriverIncentive) (models.DriverIncentive, error) {
	
	db := u.db.GetConnection()

	if err := db.WithContext(ctx).Table("driver_incentives").Save(&driversIncentive).Error; err != nil {
		return models.DriverIncentive{}, err
	}

	if err := db.WithContext(ctx).
		Preload("Booking").
		Preload("Booking.Customer.Membership").
		Preload("Booking.Customer").
		Preload("Booking.Car").
		Preload("Booking.Driver").
		Preload("Booking.BookingType").
		First(&driversIncentive, driversIncentive.ID).Error; err != nil {
		return models.DriverIncentive{}, err
	}

	return driversIncentive, nil
}
func (u *driversIncentiveQueryImpl) EditDriversIncentive(ctx context.Context, id uint64, driDriverIncentive models.DriverIncentive) (models.DriverIncentive, error) {
	db := u.db.GetConnection()
	if err := db.WithContext(ctx).
		Model(&models.DriverIncentive{}).
		Where("id = ?", id).
		Updates(driDriverIncentive).Error; err != nil {
		return models.DriverIncentive{}, err
	}

	updatedDriverIncentive := models.DriverIncentive{}
	if err := db.WithContext(ctx).
		Preload("Booking").
		Preload("Booking.Customer.Membership").
		Preload("Booking.Customer").
		Preload("Booking.Car").
		Preload("Booking.Driver").
		Preload("Booking.BookingType").
		First(&updatedDriverIncentive, id).Error; err != nil {
		return models.DriverIncentive{}, err
	}

	return updatedDriverIncentive, nil
}
