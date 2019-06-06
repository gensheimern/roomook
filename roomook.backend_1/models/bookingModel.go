package models

//Booking struct which represents de table booking in the database
type BookingModel struct {
	BookingID string   `json:"bookingID"`
	ParentID string `json:"parentID"`
	Title     string   `json:"title"`
	UserID    string   `json:"userID"`
	RoomID    string   `json:"roomID"`
	Begin     []string `json:"begin"`
	End       []string `json:"end"`
	RecurringType string `json:"recurringType"`

}

