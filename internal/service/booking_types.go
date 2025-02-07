package service

import (
	"car-rental/internal/models"
	"car-rental/internal/repository"
	"context"
	"errors"
	"time"
)

type BookingTypeservice interface {
	GetBookingTypes(ctx context.Context) ([]models.BookingType, error)
	GetBookingTypesByID(ctx context.Context, id uint64) (models.BookingType, error)
	CreateBookingType(ctx context.Context, bookingType models.InputBookingType) (models.BookingType, error)
	EditBookingType(ctx context.Context, id uint64, bookingType models.InputBookingType) (models.BookingType, error)
	DeleteBookingType(ctx context.Context, id uint64) (models.BookingType, error)
}
type bookingTypeserviceImpl struct {
	bookingTypeRepo repository.BookingTypesQuery
}

func NewBookingTypeservice(bookingTypeRepo repository.BookingTypesQuery) BookingTypeservice {
	return &bookingTypeserviceImpl{bookingTypeRepo: bookingTypeRepo}
}


func (s *bookingTypeserviceImpl) GetBookingTypes(ctx context.Context) ([]models.BookingType, error) {
	bookingTypes, err := s.bookingTypeRepo.GetBookingTypes(ctx)
	if err != nil {
		return nil, err
	}
	return bookingTypes, nil
}

func (s *bookingTypeserviceImpl) GetBookingTypesByID(ctx context.Context, id uint64) (models.BookingType, error) {
	bookingType, err := s.bookingTypeRepo.GetBookingTypesByID(ctx, id)
	if err != nil {
		return models.BookingType{}, err
	}
	if bookingType.ID == 0 {
		return models.BookingType{}, errors.New("bookingType not found")
	}
	return bookingType, nil
}

func (s *bookingTypeserviceImpl) CreateBookingType(ctx context.Context, bookingType models.InputBookingType) (models.BookingType, error) {
	NewBookingType := models.BookingType{}
	NewBookingType.BookingType = bookingType.BookingType
	NewBookingType.Description = bookingType.Description
	NewBookingType.CreatedAt = time.Now()

	createdBookingType, err := s.bookingTypeRepo.CreateBookingTypes(ctx, NewBookingType)
	if err != nil {
		return models.BookingType{}, err
	}
	return createdBookingType, nil
}

func (s *bookingTypeserviceImpl) EditBookingType(ctx context.Context, id uint64, bookingType models.InputBookingType) (models.BookingType, error) {
	updatedBookingType := models.BookingType{}
	updatedBookingType.BookingType = bookingType.BookingType
	updatedBookingType.Description = bookingType.Description
	updatedBookingType.UpdatedAt = time.Now()

	updatedBookingType, err := s.bookingTypeRepo.EditBookingTypes(ctx, id, updatedBookingType)
	if err != nil {
		return models.BookingType{}, err
	}
	return updatedBookingType, nil
}

func (s *bookingTypeserviceImpl) DeleteBookingType(ctx context.Context, id uint64) (models.BookingType, error) {
	bookingType, err := s.bookingTypeRepo.GetBookingTypesByID(ctx, id)
	if err != nil {
		return models.BookingType{}, err
	}
	if bookingType.ID == 0 {
		return models.BookingType{}, nil
	}

	err = s.bookingTypeRepo.DeleteBookingTypesByID(ctx, id)
	if err != nil {
		return models.BookingType{}, err
	}

	return bookingType, err
}
