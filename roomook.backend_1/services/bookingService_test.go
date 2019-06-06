package services_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mocks "jacob.de/roomook/roomook.backend/interfaces/mocks"
	"jacob.de/roomook/roomook.backend/models"
	"jacob.de/roomook/roomook.backend/services"
)

func TestBookingExists(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	bookingRepository := mocks.NewMockIBookingRepository(mockCtrl)

	bookingRepository.EXPECT().BookingExists(gomock.Eq("33")).Return(true, nil)

	mockInjection := &services.BookingService{Repo: bookingRepository}
	exists, _ := mockInjection.SBookingExists("33")

	assert.Equal(t, true, exists)

}

func TestBookingExistsByDate(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	bookingRepository := mocks.NewMockIBookingRepository(mockCtrl)

	bookingRepository.EXPECT().BookingExistsByDateTime("2019-05-01 12:00", "2019-05-01 13:00").Return(false, nil)

	mockInjection := &services.BookingService{Repo: bookingRepository}
	exists, _ := mockInjection.SBookingExistsByDateTime("2019-05-01 12:00", "2019-05-01 13:00")

	assert.Equal(t, false, exists)

}

func TestGetAllBooking(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	var mockModelSlice []models.BookingModel

	mockModelSlice = append(mockModelSlice,

		models.BookingModel{
			BookingID: "233",
			UserID:    "nico@jacob.de",
			RoomID:    "22A",
			Begin:     "2019-05-01 12:00",
			End:       "2019-05-01 13:00"},
		models.BookingModel{
			BookingID: "444",
			UserID:    "hans@jacob.de",
			RoomID:    "13B",
			Begin:     "2019-05-01 14:00",
			End:       "2019-05-01 16:00"})

	bookingRepository := mocks.NewMockIBookingRepository(mockCtrl)

	bookingRepository.EXPECT().GetAllBooking().Return(mockModelSlice, nil)

	mockInjection := &services.BookingService{Repo: bookingRepository}
	out, _ := mockInjection.SGetAllBooking()

	var expectedModelSlice []models.BookingModel

	expectedModelSlice = append(expectedModelSlice,

		models.BookingModel{
			BookingID: "233",
			UserID:    "nico@jacob.de",
			RoomID:    "22A",
			Begin:     "2019-05-01 12:00",
			End:       "2019-05-01 13:00"},
		models.BookingModel{
			BookingID: "444",
			UserID:    "hans@jacob.de",
			RoomID:    "13B",
			Begin:     "2019-05-01 14:00",
			End:       "2019-05-01 16:00"})

	assert.Equal(t, expectedModelSlice, out)
}

func TestGetBookingByID(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	bookingRepository := mocks.NewMockIBookingRepository(mockCtrl)

	callFirst := bookingRepository.EXPECT().BookingExists("233").Return(true, nil)

	bookingRepository.EXPECT().GetBookingByID("233").Return(models.BookingModel{
		BookingID: "233",
		UserID:    "nico@jacob.de",
		RoomID:    "22A",
		Begin:     "2019-05-01 12:00",
		End:       "2019-05-01 13:00"}, nil).After(callFirst)

	mockInjection := &services.BookingService{Repo: bookingRepository}
	out, _ := mockInjection.SGetBookingByID("233")

	assert.Equal(t, models.BookingModel{
		BookingID: "233",
		UserID:    "nico@jacob.de",
		RoomID:    "22A",
		Begin:     "2019-05-01 12:00",
		End:       "2019-05-01 13:00"}, out)

}

func TestCreateBooking(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	bookingRepository := mocks.NewMockIBookingRepository(mockCtrl)

	booking := models.BookingModel{
		BookingID: "233",
		UserID:    "nico@jacob.de",
		RoomID:    "22A",
		Begin:     "2019-05-01 12:00",
		End:       "2019-05-01 13:00",
	}
	bookingRepository.EXPECT().CreateBooking(booking).Return(booking, nil)

	mockInjection := &services.BookingService{Repo: bookingRepository}
	out, _ := mockInjection.SCreateBooking(booking)

	assert.Equal(t, models.BookingModel{
		BookingID: "233",
		UserID:    "nico@jacob.de",
		RoomID:    "22A",
		Begin:     "2019-05-01 12:00",
		End:       "2019-05-01 13:00",
	}, out)

}

func TestUpdateBooking(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	bookingRepository := mocks.NewMockIBookingRepository(mockCtrl)

	booking := models.BookingModel{
		BookingID: "233",
		UserID:    "nico@jacob.de",
		RoomID:    "22A",
		Begin:     "2019-05-01 12:00",
		End:       "2019-05-01 13:00",
	}

	callFirst := bookingRepository.EXPECT().BookingExists("4044").Return(true, nil)

	bookingRepository.EXPECT().UpdateBooking("4044", booking).Return(booking, nil).After(callFirst)

	mockInjection := &services.BookingService{Repo: bookingRepository}
	out, _ := mockInjection.SUpdateBooking("4044", booking)
	assert.Equal(t, models.BookingModel{
		BookingID: "233",
		UserID:    "nico@jacob.de",
		RoomID:    "22A",
		Begin:     "2019-05-01 12:00",
		End:       "2019-05-01 13:00",
	}, out)
}

func TestDeleteBooking(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	bookingRepository := mocks.NewMockIBookingRepository(mockCtrl)

	callFirst := bookingRepository.EXPECT().BookingExists("233").Return(true, nil)

	bookingRepository.EXPECT().DeleteBooking("233").Return(models.BookingModel{
		BookingID: "233",
		UserID:    "nico@jacob.de",
		RoomID:    "22A",
		Begin:     "2019-05-01 12:00",
		End:       "2019-05-01 13:00",
	}, nil).After(callFirst)

	mockInjection := &services.BookingService{Repo: bookingRepository}
	out, _ := mockInjection.SDeleteBooking("233")

	assert.Equal(t, models.BookingModel{
		BookingID: "233",
		UserID:    "nico@jacob.de",
		RoomID:    "22A",
		Begin:     "2019-05-01 12:00",
		End:       "2019-05-01 13:00",
	}, out)

}
