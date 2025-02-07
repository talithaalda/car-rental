package repository

import (
	"context"

	"car-rental/internal/infrastructure"
	"car-rental/internal/models"

	"gorm.io/gorm"
)

type DriversQuery interface {
	GetDrivers(ctx context.Context) ([]models.Driver, error)
	GetDriversByID(ctx context.Context, id uint64) (models.Driver, error)
	EditDrivers(ctx context.Context, id uint64, drivers models.Driver) (models.Driver, error)
	DeleteDriversByID(ctx context.Context, id uint64) error
	CreateDrivers(ctx context.Context, drivers models.Driver) (models.Driver, error)
}

type DriversCommand interface {
	CreateDrivers(ctx context.Context, drivers models.Driver) (models.Driver, error)
}

type driversQueryImpl struct {
	db infrastructure.GormPostgres
}

func NewDriversQuery(db infrastructure.GormPostgres) DriversQuery {
	return &driversQueryImpl{db: db}
}

func (u *driversQueryImpl) GetDrivers(ctx context.Context) ([]models.Driver, error) {
	db := u.db.GetConnection()
	drivers := []models.Driver{}
	if err := db.
		WithContext(ctx).
		Table("drivers").
		Find(&drivers).Error; err != nil {
		return nil, err
	}
	return drivers, nil
}

func (u *driversQueryImpl) GetDriversByID(ctx context.Context, id uint64) (models.Driver, error) {
	db := u.db.GetConnection()
	drivers := models.Driver{}
	if err := db.
		WithContext(ctx).
		Table("drivers").
		Where("id = ?", id).
		Find(&drivers).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Driver{}, nil
		}

		return models.Driver{}, err
	}
	return drivers, nil
}

func (u *driversQueryImpl) DeleteDriversByID(ctx context.Context, id uint64) error {
	db := u.db.GetConnection()
	if err := db.
		WithContext(ctx).
		Table("drivers").
		Delete(&models.Driver{ID: uint(id)}).
		Error; err != nil {
		return err
	}
	return nil
}

func (u *driversQueryImpl) CreateDrivers(ctx context.Context, drivers models.Driver) (models.Driver, error) {
	
	db := u.db.GetConnection()
	if err := db.
		WithContext(ctx).
		Table("drivers").
		Save(&drivers).Error; err != nil {
		return models.Driver{}, err
	}
	return drivers, nil
}
func (u *driversQueryImpl) EditDrivers(ctx context.Context, id uint64, driver models.Driver) (models.Driver, error) {
	db := u.db.GetConnection()
	updatedDrivers := models.Driver{}
	if err := db.
		WithContext(ctx).
		Table("drivers").
		Where("id = ?", id).Updates(&driver).First(&updatedDrivers).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return models.Driver{}, nil
			}
		}
	return updatedDrivers, nil
}
