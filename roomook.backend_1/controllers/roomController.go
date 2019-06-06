package controllers

import (
	"net/http"
	"fmt"
	"jacob.de/roomook/roomook.backend/models"

	"github.com/gin-gonic/gin"
	"jacob.de/roomook/roomook.backend/interfaces"
)

type RoomController struct {
	interfaces.IRoomService
}

func (controller *RoomController) GetAllRoom(ctx *gin.Context) {

	room, err := controller.SGetAllRoom()
	if err != nil {

		errMsg := err.Error()
		ctx.JSON(http.StatusBadRequest, models.ErrorMSG{Error: errMsg})
	} else {
		ctx.JSONP(http.StatusOK, room)
	}
}

func (controller *RoomController) GetRoomByID(ctx *gin.Context) {

	roomID := ctx.Params.ByName("id")

	room, err := controller.SGetRoomByID(roomID)
	if err != nil {

		errMsg := err.Error()
		ctx.JSON(http.StatusBadRequest, models.ErrorMSG{Error: errMsg})

	} else {
		ctx.JSONP(http.StatusOK, room)
	}
}

func (controller *RoomController) GetRoomByDate(ctx *gin.Context) {

	type TmpDate struct {
		Begin string `json:"begin"`
		End string`json:"end"`
	}

	type TmpRommByDate struct {
		Begin string `json:"begin"`
		End string`json:"end"`
		Rooms []models.RoomModel `json:"rooms"`
	}

	var tmpDate TmpDate
	ctx.Bind(&tmpDate)	
	fmt.Println("hier kommts tmpdate")

fmt.Println(tmpDate)

	room, begin, end, err := controller.SGetRoomByDate(tmpDate.Begin, tmpDate.End)
	if err != nil {

		errMsg := err.Error()
		ctx.JSON(http.StatusBadRequest, models.ErrorMSG{Error: errMsg})
		
	
	} else {
		roomByDate := TmpRommByDate{
			Begin: begin,
			End: end,
			Rooms: room,
		}
		ctx.JSONP(http.StatusOK, roomByDate)
	}
}

func (controller *RoomController) CreateRoom(ctx *gin.Context) {

	var room models.RoomModel
	ctx.Bind(&room)

	room, err := controller.SCreateRoom(room)
	if err != nil {

		errMsg := err.Error()
		ctx.JSON(http.StatusBadRequest, models.ErrorMSG{Error: errMsg})

	} else {
		ctx.JSONP(http.StatusCreated, room)
	}
}

func (controller *RoomController) UpdateRoom(ctx *gin.Context) {

	roomID := ctx.Params.ByName("id")

	var roomToUpdate models.RoomModel
	ctx.Bind(&roomToUpdate)

	room, err := controller.SUpdateRoom(roomID, roomToUpdate)
	if err != nil {

		errMsg := err.Error()
		ctx.JSON(http.StatusBadRequest, models.ErrorMSG{Error: errMsg})

	} else {
		ctx.JSONP(http.StatusOK, room)
	}
}

func (controller *RoomController) DeleteRoom(ctx *gin.Context) {

	roomID := ctx.Params.ByName("id")

	room, err := controller.SDeleteRoom(roomID)
	if err != nil {

		errMsg := err.Error()
		ctx.JSON(http.StatusBadRequest, models.ErrorMSG{Error: errMsg})

	} else {
		ctx.JSONP(http.StatusOK, room)
	}
}
