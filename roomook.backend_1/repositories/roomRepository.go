package repositories

import (
	"errors"
"fmt"
	"jacob.de/roomook/roomook.backend/interfaces"
	"jacob.de/roomook/roomook.backend/models"
)

// TODO put down some Comments

type RoomRepository struct {
	interfaces.IDbHandler
}

func (repositorys *RoomRepository) RoomExists(roomID string) (bool, error) {

	rows, err := repositorys.PreparedQuery("SELECT count(*) FROM room WHERE RoomID = ?", roomID)
	defer rows.Close()

	if err != nil {
		return false, err
	}

	var exists bool
	rows.Next()
	rows.Scan(&exists)

	if exists != true {
		err := errors.New("room not found")
		return exists, err
	}

	return exists, err
}

//GetAllRoom sdfsfdsf
func (repositorys *RoomRepository) GetAllRoom() ([]models.RoomModel, error) {

	rows, err := repositorys.Query("SELECT * FROM room")
	defer rows.Close()

	if err != nil {
		return []models.RoomModel{}, err
	}

	var rooms []models.RoomModel
	var room models.RoomModel

	for rows.Next() {
		rows.Scan(&room.RoomID, &room.Name)
		rooms = append(rooms, room)
	}

	return rooms, err
}

func (repositorys *RoomRepository) GetRoomByID(roomID string) (models.RoomModel, error) {

	rows, err := repositorys.PreparedQuery("SELECT * FROM room WHERE RoomID = ?", roomID)
	defer rows.Close()

	if err != nil {
		return models.RoomModel{}, err
	}

	var room models.RoomModel
	rows.Next()
	rows.Scan(&room.RoomID, &room.Name)

	return room, err
}

func (repositorys *RoomRepository) GetRoomByDate(begin string, end string) ([]models.RoomModel, error) {

	fmt.Println(begin, end)
	rows, err := repositorys.PreparedQuery(
		"SELECT * FROM room WHERE RoomID NOT IN	(SELECT RoomID from booking WHERE   ((Begin BETWEEN ? AND ?) OR (End BETWEEN ? AND ?) OR (Begin <= ? AND End >= ?)))",
		begin, end, begin, end, begin, end )
	defer rows.Close()

	if err != nil {
		return []models.RoomModel{}, err
	}
	var rooms []models.RoomModel
	var room models.RoomModel
	for rows.Next(){
		rows.Scan(&room.RoomID, &room.Name)
		rooms = append(rooms,room)
	}
	return rooms, err
}

func (repositorys *RoomRepository) CreateRoom(room models.RoomModel) (models.RoomModel, error) {

	_, err := repositorys.PreparedQuery("INSERT INTO room (RoomID, Name) VALUES( ?, ?)",

		room.RoomID, room.Name)

	if err != nil {
		return models.RoomModel{}, err
	}

	return room, err
}

func (repositorys *RoomRepository) UpdateRoom(roomID string, roomToUpdate models.RoomModel) (models.RoomModel, error) {

	rows, err := repositorys.PreparedQuery("UPDATE room set RoomID= ?, Name= ? where RoomID= ?;",
		roomToUpdate.RoomID,
		roomToUpdate.Name,
		roomID)
	defer rows.Close()

	if err != nil {
		return models.RoomModel{}, err
	}

	var room models.RoomModel
	// FIXME should i send back what the DB actual gives back or should i give back the object
	rows.Next()
	rows.Scan(&room.RoomID, &room.Name)

	return room, err
}

func (repositorys *RoomRepository) DeleteRoom(roomID string) (models.RoomModel, error) {

	rows, err := repositorys.PreparedQuery("DELETE FROM room WHERE RoomID = ?", roomID)
	defer rows.Close()

	if err != nil {
		return models.RoomModel{}, err
	}

	var room models.RoomModel
	rows.Next()
	rows.Scan(&room.RoomID, &room.Name)

	return room, err
}
