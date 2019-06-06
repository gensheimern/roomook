package interfaces

//go:generate mockgen -destination=./mocks/IRoomRepository.go -package=mocks jacob.de/roombook/interfaces IRoomRepository

import (
	"jacob.de/roomook/roomook.backend/models"
)

//IRoomRepository Interface
type IRoomRepository interface {
	RoomExists(roomID string) (bool, error)
	GetAllRoom() ([]models.RoomModel, error)
	GetRoomByID(roomID string) (models.RoomModel, error)
	GetRoomByDate(begin string, end string) ([]models.RoomModel,  error)
	CreateRoom(room models.RoomModel) (models.RoomModel, error)
	UpdateRoom(roomID string, roomToUpdate models.RoomModel) (models.RoomModel, error)
	DeleteRoom(roomID string) (models.RoomModel, error)
}
