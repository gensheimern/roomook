package services

import (
	"fmt"

	"jacob.de/roomook/roomook.backend/models"

	"jacob.de/roomook/roomook.backend/interfaces"
)

type BookingService struct {
	Repo interfaces.IBookingRepository
}

func (service *BookingService) SBookingExists(bookingID string) (bool, error) {

	exists, err := service.Repo.BookingExists(bookingID)
	return exists, err

}

func (service *BookingService) SBookingExistsByDateTime(roomID string, begin string, end string) (bool, error) {

	exists, err := service.Repo.BookingExistsByDateTime(roomID, begin, end)
	return exists, err

}

func (service *BookingService) SGetAllBooking() ([]models.BookingModel, error) {

	booking, err := service.Repo.GetAllBooking()
	if err != nil {
		return booking, err
	}

	return booking, nil
}

func (service *BookingService) SGetAllBookingByRoomID(roomID string) ([]models.BookingModel, error) {

	booking, err := service.Repo.GetAllBookingByRoomID(roomID)
	if err != nil {
		return booking, err
	}

	return booking, nil
}

func (service *BookingService) SGetBookingByID(bookingID string) (models.BookingModel, error) {

	exists, err := service.Repo.BookingExists(bookingID)

	if exists {

		booking, err := service.Repo.GetBookingByID(bookingID)
		if err != nil {

			return booking, nil
		}

		return booking, err
	}

	return models.BookingModel{}, err

}

func (service *BookingService) SCreateBooking(booking models.BookingModel) ([]models.BookingModel, error) {

	for index, _ := range booking.Begin {
		exists, err := service.Repo.BookingExistsByDateTime(booking.RoomID, booking.Begin[index], booking.End[index])

		if exists {
			return []models.BookingModel{}, err
		}
	}
	bookings, err := service.Repo.CreateBooking(booking)

	if err != nil {
		fmt.Println(err)
		return bookings, err
	}
	return bookings, nil
}

func (service *BookingService) SUpdateBooking(bookingID string, bookingToUpdate models.BookingModel) (models.BookingModel, error) {

	exists, err := service.Repo.BookingExists(bookingID)

	if exists {
		booking, err := service.Repo.UpdateBooking(bookingID, bookingToUpdate)
		if err != nil {

			return booking, nil
		}

		return booking, err
	}
	return models.BookingModel{}, err

}

func (service *BookingService) SDeleteBooking(bookingID string) (models.BookingModel, error) {
	exists, err := service.Repo.BookingExists(bookingID)

	if exists {
		booking, err := service.Repo.DeleteBooking(bookingID)
		if err != nil {

			return booking, nil
		}

		return booking, err
	}
	return models.BookingModel{}, err

}

func (service *BookingService) SDeleteRecurringBooking(parentID string) (models.BookingModel, error) {
	exists, err := service.Repo.BookingParentExists(parentID)

	if exists {
		booking, err := service.Repo.DeleteRecurringBooking(parentID)
		if err != nil {

			return booking, nil
		}

		return booking, err
	}
	return models.BookingModel{}, err

}
