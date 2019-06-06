package interfaces

//go:generate mockgen -destination=./mocks/IRoomService.go -package=mocks jacob.de/roombook/interfaces IRoomService

import (
	"jacob.de/roomook/roomook.backend/models"
)

type IRoomService interface {
	SRoomExists(roomID string) (bool, error)
	SGetAllRoom() ([]models.RoomModel, error)
	SGetRoomByID(roomID string) (models.RoomModel, error)
	SGetRoomByDate(begin string, end string) ([]models.RoomModel,  string,  string, error)
	SCreateRoom(room models.RoomModel) (models.RoomModel, error)
	SUpdateRoom(roomID string, roomToUpdate models.RoomModel) (models.RoomModel, error)
	SDeleteRoom(roomID string) (models.RoomModel, error)
}
