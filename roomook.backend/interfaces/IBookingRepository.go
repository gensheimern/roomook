package interfaces

//go:generate mockgen -destination=./mocks/IBookingRepository.go -package=mocks jacob.de/roombook/interfaces IBookingRepository

import (
	"jacob.de/roomook/roomook.backend/models"
)

//IBookingRepository Interface
type IBookingRepository interface {
	BookingExists(bookingID string) (bool, error)
	BookingParentExists(parentID string) (bool, error)
	BookingExistsByDateTime(roomID string, begin string, end string) (bool, error)
	GetAllBooking() ([]models.BookingModel, error)
	GetAllBookingByRoomID(roomID string) ([]models.BookingModel, error)
	GetBookingByID(bookingID string) (models.BookingModel, error)
	CreateBooking(booking models.BookingModel) ([]models.BookingModel, error)
	UpdateBooking(bookingID string, bookingToUpdate models.BookingModel) (models.BookingModel, error)
	DeleteBooking(bookingID string) (models.BookingModel, error)
	DeleteRecurringBooking(parentID string) (models.BookingModel, error)

}
