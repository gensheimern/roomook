package services

import (
	"jacob.de/roomook/roomook.backend/models"

	"jacob.de/roomook/roomook.backend/interfaces"
)

type RoomService struct {
	Repo interfaces.IRoomRepository
}

func (service *RoomService) SRoomExists(roomID string) (bool, error) {

	exists, err := service.Repo.RoomExists(roomID)

	return exists, err

}

func (service *RoomService) SGetAllRoom() ([]models.RoomModel, error) {
	room, err := service.Repo.GetAllRoom()
	if err != nil {
		return room, err
	}

	return room, nil
}

func (service *RoomService) SGetRoomByID(roomID string) (models.RoomModel, error) {
	exists, err := service.Repo.RoomExists(roomID)

	if exists {
		room, err := service.Repo.GetRoomByID(roomID)
		if err != nil {

			return room, err
		}

		return room, nil
	}
	return models.RoomModel{}, err

}

func (service *RoomService) SGetRoomByDate(begin string, end string) ([]models.RoomModel, string, string, error) {


		room, err := service.Repo.GetRoomByDate(begin, end)
		if err != nil {

			return room, begin, end, err
		}

		
	
	return room, begin, end, err

}

func (service *RoomService) SCreateRoom(room models.RoomModel) (models.RoomModel, error) {

	room, err := service.Repo.CreateRoom(room)
	if err != nil {
		return room, err
	}

	return room, nil
}

func (service *RoomService) SUpdateRoom(roomID string, roomToUpdate models.RoomModel) (models.RoomModel, error) {

	exists, err := service.Repo.RoomExists(roomID)

	if exists {
		room, err := service.Repo.UpdateRoom(roomID, roomToUpdate)
		if err != nil {

			return room, nil
		}

		return room, err
	}
	return models.RoomModel{}, err

}

func (service *RoomService) SDeleteRoom(roomID string) (models.RoomModel, error) {

	exists, err := service.Repo.RoomExists(roomID)

	if exists {
		room, err := service.Repo.DeleteRoom(roomID)
		if err != nil {

			return room, nil
		}

		return room, err
	}
	return models.RoomModel{}, err

}
