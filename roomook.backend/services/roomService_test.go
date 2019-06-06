package services_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mocks "jacob.de/roomook/roomook.backend/interfaces/mocks"
	"jacob.de/roomook/roomook.backend/models"
	"jacob.de/roomook/roomook.backend/services"
)

func TestRoomExists(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	roomRepository := mocks.NewMockIRoomRepository(mockCtrl)

	roomRepository.EXPECT().RoomExists(gomock.Eq("33A")).Return(true, nil)

	mockInjection := &services.RoomService{Repo: roomRepository}
	exists, _ := mockInjection.SRoomExists("33A")

	assert.Equal(t, true, exists)

}

func TestGetAllRoom(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	var mockModelSlice []models.RoomModel

	mockModelSlice = append(mockModelSlice,

		models.RoomModel{
			RoomID: "33A",
			Name:   "Aquarium",
		},
		models.RoomModel{
			RoomID: "444",
			Name:   "Hutraum"})

	roomRepository := mocks.NewMockIRoomRepository(mockCtrl)

	roomRepository.EXPECT().GetAllRoom().Return(mockModelSlice, nil)

	mockInjection := &services.RoomService{Repo: roomRepository}
	out, _ := mockInjection.SGetAllRoom()

	var expectedModelSlice []models.RoomModel

	expectedModelSlice = append(expectedModelSlice,

		models.RoomModel{
			RoomID: "33A",
			Name:   "Aquarium",
		},
		models.RoomModel{
			RoomID: "444",
			Name:   "Hutraum"})

	assert.Equal(t, expectedModelSlice, out)
}

func TestGetRoomByID(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	roomRepository := mocks.NewMockIRoomRepository(mockCtrl)

	callFirst := roomRepository.EXPECT().RoomExists("33A").Return(true, nil)

	roomRepository.EXPECT().GetRoomByID("33A").Return(models.RoomModel{
		RoomID: "33A",
		Name:   "Aquarium",
	}, nil).After(callFirst)

	mockInjection := &services.RoomService{Repo: roomRepository}
	out, _ := mockInjection.SGetRoomByID("33A")

	assert.Equal(t, models.RoomModel{
		RoomID: "33A",
		Name:   "Aquarium",
	}, out)

}

func TestCreateRoom(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	roomRepository := mocks.NewMockIRoomRepository(mockCtrl)

	room := models.RoomModel{
		RoomID: "33A",
		Name:   "Aquarium",
	}

	roomRepository.EXPECT().CreateRoom(room).Return(room, nil)

	mockInjection := &services.RoomService{Repo: roomRepository}
	out, _ := mockInjection.SCreateRoom(room)

	assert.Equal(t, models.RoomModel{
		RoomID: "33A",
		Name:   "Aquarium",
	}, out)

}

func TestUpdateRoom(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	roomRepository := mocks.NewMockIRoomRepository(mockCtrl)

	room := models.RoomModel{
		RoomID: "33A",
		Name:   "Aquarium",
	}

	callFirst := roomRepository.EXPECT().RoomExists("4044").Return(true, nil)

	roomRepository.EXPECT().UpdateRoom("4044", room).Return(room, nil).After(callFirst)

	mockInjection := &services.RoomService{Repo: roomRepository}
	out, _ := mockInjection.SUpdateRoom("4044", room)
	assert.Equal(t, models.RoomModel{
		RoomID: "33A",
		Name:   "Aquarium",
	}, out)
}

func TestDeleteRoom(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	roomRepository := mocks.NewMockIRoomRepository(mockCtrl)

	callFirst := roomRepository.EXPECT().RoomExists("33A").Return(true, nil)

	roomRepository.EXPECT().DeleteRoom("33A").Return(models.RoomModel{
		RoomID: "33A",
		Name:   "Aquarium",
	}, nil).After(callFirst)

	mockInjection := &services.RoomService{Repo: roomRepository}
	out, _ := mockInjection.SDeleteRoom("33A")

	assert.Equal(t, models.RoomModel{
		RoomID: "33A",
		Name:   "Aquarium",
	}, out)

}
