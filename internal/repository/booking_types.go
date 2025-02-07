package repository

import (
	"context"

	"car-rental/internal/infrastructure"
	"car-rental/internal/models"

	"gorm.io/gorm"
)

type BookingTypesQuery interface {
	GetBookingTypes(ctx context.Context) ([]models.BookingType, error)
	GetBookingTypesByID(ctx context.Context, id uint64) (models.BookingType, error)
	EditBookingTypes(ctx context.Context, id uint64, bookingTypes models.BookingType) (models.BookingType, error)
	DeleteBookingTypesByID(ctx context.Context, id uint64) error
	CreateBookingTypes(ctx context.Context, bookingTypes models.BookingType) (models.BookingType, error)
}

type BookingTypesCommand interface {
	CreateBookingTypes(ctx context.Context, bookingTypes models.BookingType) (models.BookingType, error)
}

type bookingTypesQueryImpl struct {
	db infrastructure.GormPostgres
}

func NewBookingTypesQuery(db infrastructure.GormPostgres) BookingTypesQuery {
	return &bookingTypesQueryImpl{db: db}
}

func (u *bookingTypesQueryImpl) GetBookingTypes(ctx context.Context) ([]models.BookingType, error) {
	db := u.db.GetConnection()
	bookingTypes := []models.BookingType{}
	if err := db.
		WithContext(ctx).
		Table("booking_types").
		Find(&bookingTypes).Error; err != nil {
		return nil, err
	}
	return bookingTypes, nil
}

func (u *bookingTypesQueryImpl) GetBookingTypesByID(ctx context.Context, id uint64) (models.BookingType, error) {
	db := u.db.GetConnection()
	bookingTypes := models.BookingType{}
	if err := db.
		WithContext(ctx).
		Table("booking_types").
		Where("id = ?", id).
		Find(&bookingTypes).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.BookingType{}, nil
		}

		return models.BookingType{}, err
	}
	return bookingTypes, nil
}

func (u *bookingTypesQueryImpl) DeleteBookingTypesByID(ctx context.Context, id uint64) error {
	db := u.db.GetConnection()
	if err := db.
		WithContext(ctx).
		Table("booking_types").
		Delete(&models.BookingType{ID: uint(id)}).
		Error; err != nil {
		return err
	}
	return nil
}

func (u *bookingTypesQueryImpl) CreateBookingTypes(ctx context.Context, bookingTypes models.BookingType) (models.BookingType, error) {
	
	db := u.db.GetConnection()
	if err := db.
		WithContext(ctx).
		Table("booking_types").
		Save(&bookingTypes).Error; err != nil {
		return models.BookingType{}, err
	}
	return bookingTypes, nil
}
func (u *bookingTypesQueryImpl) EditBookingTypes(ctx context.Context, id uint64, bookingType models.BookingType) (models.BookingType, error) {
	db := u.db.GetConnection()
	updatedBookingTypes := models.BookingType{}
	if err := db.
		WithContext(ctx).
		Table("booking_types").
		Where("id = ?", id).Updates(&bookingType).First(&updatedBookingTypes).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return models.BookingType{}, nil
			}
		}
	return updatedBookingTypes, nil
}
