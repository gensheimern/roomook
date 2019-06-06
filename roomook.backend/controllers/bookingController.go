package controllers

import (
	"net/http"
	"fmt"
	"jacob.de/roomook/roomook.backend/models"

	"github.com/gin-gonic/gin"
	"jacob.de/roomook/roomook.backend/interfaces"
)

type BookingController struct {
	interfaces.IBookingService
}

func (controller *BookingController) GetAllBooking(ctx *gin.Context) {
	booking, err := controller.SGetAllBooking()
	if err != nil {
		
		errMsg := err.Error()
		ctx.JSON(http.StatusBadRequest, models.ErrorMSG{Error: errMsg})

	} else {
		ctx.JSONP(http.StatusOK, booking)
	}
}

func (controller *BookingController) GetAllBookingByRoomID(ctx *gin.Context) {

	roomID := ctx.Params.ByName("id")

	booking, err := controller.SGetAllBookingByRoomID(roomID)
	if err != nil {

		errMsg := err.Error()
		ctx.JSON(http.StatusBadRequest, models.ErrorMSG{Error: errMsg})

	} else {
		ctx.JSONP(http.StatusOK, booking)
	}
}

func (controller *BookingController) GetBookingByID(ctx *gin.Context) {

	bookingID := ctx.Params.ByName("id")

	booking, err := controller.SGetBookingByID(bookingID)
	if err != nil {

		errMsg := err.Error()
		ctx.JSON(http.StatusBadRequest, models.ErrorMSG{Error: errMsg})

	} else {
		ctx.JSONP(http.StatusOK, booking)
	}
}

func (controller *BookingController) CreateBooking(ctx *gin.Context) {
	
	var booking models.BookingModel
	ctx.Bind(&booking)
	booking.UserID =  ctx.Request.Header.Get("x-authorized-user")
	fmt.Println(booking)
	var bookings []models.BookingModel
	bookings, err := controller.SCreateBooking(booking)
	if err != nil {

		errMsg := err.Error()
		ctx.JSON(http.StatusBadRequest, models.ErrorMSG{Error: errMsg})

	} else {
		ctx.JSON(http.StatusCreated, bookings)
	}
}

func (controller *BookingController) UpdateBooking(ctx *gin.Context) {

	bookingID := ctx.Params.ByName("id")

	var bookingToUpdate models.BookingModel
	ctx.Bind(&bookingToUpdate)

	booking, err := controller.SUpdateBooking(bookingID, bookingToUpdate)
	if err != nil {

		errMsg := err.Error()
		ctx.JSON(http.StatusBadRequest, models.ErrorMSG{Error: errMsg})

	} else {
		ctx.JSONP(http.StatusOK, booking)
	}
}

func (controller *BookingController) DeleteBooking(ctx *gin.Context) {

	bookingID := ctx.Params.ByName("id")

	booking, err := controller.SDeleteBooking(bookingID)
	if err != nil {

		errMsg := err.Error()
		ctx.JSON(http.StatusBadRequest, models.ErrorMSG{Error: errMsg})

	} else {
		ctx.JSONP(http.StatusOK, booking)
	}
}

func (controller *BookingController) DeleteRecurringBooking(ctx *gin.Context) {
	parrentID := ctx.Params.ByName("id")

	booking, err := controller.SDeleteRecurringBooking(parrentID)
	if err != nil {

		errMsg := err.Error()
		ctx.JSON(http.StatusBadRequest, models.ErrorMSG{Error: errMsg})

	} else {
		ctx.JSONP(http.StatusOK, booking)
	}
}
