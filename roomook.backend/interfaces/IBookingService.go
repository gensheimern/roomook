package interfaces

//go:generate mockgen -destination=./mocks/IBookingService.go -package=mocks jacob.de/roombook/interfaces IBookingService

import (
	"jacob.de/roomook/roomook.backend/models"
)

type IBookingService interface {
	SBookingExists(bookingID string) (bool, error)
	SBookingExistsByDateTime(roomID string, begin string, end string) (bool, error)
	SGetAllBooking() ([]models.BookingModel, error)
	SGetAllBookingByRoomID(roomID string) ([]models.BookingModel, error)
	SGetBookingByID(bookingID string) (models.BookingModel, error)
	SCreateBooking(booking models.BookingModel) ([]models.BookingModel, error)
	SUpdateBooking(bookingID string, bookingToUpdate models.BookingModel) (models.BookingModel, error)
	SDeleteBooking(bookingID string) (models.BookingModel, error)
	SDeleteRecurringBooking(parentID string) (models.BookingModel, error)

}
