package models

//User struct which represents de table user in the database
type RoomModel struct {
	RoomID string `json:"roomID"`
	Name   string `json:"name"`
}
