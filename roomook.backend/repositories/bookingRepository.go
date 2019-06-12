package repositories

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"jacob.de/roomook/roomook.backend/interfaces"
	"jacob.de/roomook/roomook.backend/models"
)

// TODO put down some Comments

type BookingRepository struct {
	DBHandler interfaces.IDbHandler
}


func (repositorys *BookingRepository) BookingExists(bookingID string) (bool, error) {

	rows, err := repositorys.DBHandler.PreparedQuery(
		"SELECT count(*) from booking where BookingID =?", bookingID)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	var exists bool
	rows.Next()
	rows.Scan(&exists)

	if exists != true {
		err := errors.New("booking not found")
		return exists, err
	}

	return exists, err
}

func (repositorys *BookingRepository) BookingParentExists(parentID string) (bool, error) {

	rows, err := repositorys.DBHandler.PreparedQuery(
		"SELECT count(*) from booking_info where ParentID =?", parentID)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	var exists bool
	rows.Next()
	rows.Scan(&exists)

	if exists != true {
		err := errors.New("parent not found")
		return exists, err
	}

	return exists, err
}

func (repositorys *BookingRepository) BookingExistsByDateTime(roomID string, begin string, end string) (bool, error) {

	rows, err := repositorys.DBHandler.PreparedQuery(
		"SELECT count(*) from booking WHERE RoomID = ? AND ((Begin BETWEEN ? AND ?) OR (End BETWEEN ? AND ?) OR (Begin <= ? AND End >= ?)) ", roomID, begin, end, begin, end, begin, end)
	defer rows.Close()

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	var rowsNumber int
	var exists bool
	rows.Next()
	rows.Scan(&rowsNumber)
	if rowsNumber == 0 {
		exists = false
	} else {
		exists = true
	}

	if exists {
		err := errors.New("Invalid time from " + begin + " to " + end + ". Room already blocked")
		return exists, err
	}

	return exists, err
}

//GetAllBooking sdfsfdsf
func (repositorys *BookingRepository) GetAllBooking() ([]models.BookingModel, error) {
	rows, err := repositorys.DBHandler.Query("SELECT booking.*, RecurringType FROM booking, booking_info WHERE booking.ParentID = booking_info.ParentID")
	defer rows.Close()

	if err != nil {
		return []models.BookingModel{}, err
	}

	var bookings []models.BookingModel
	var booking models.BookingModel
	for rows.Next() {
		var tmpBegin, tmpEnd string
		rows.Scan(&booking.BookingID, &booking.ParentID, &booking.Title, &booking.RoomID, &booking.UserID, &tmpBegin, &tmpEnd, &booking.RecurringType)
		
		booking.Begin = nil
		booking.End = nil
		booking.Begin = append(booking.Begin, tmpBegin)
		booking.End = append(booking.End, tmpEnd)
	
		bookings = append(bookings, booking)
	}

	return bookings, err
}

//GetAllBooking sdfsfdsf
func (repositorys *BookingRepository) GetAllBookingByRoomID(roomID string) ([]models.BookingModel, error) {
	rows, err := repositorys.DBHandler.PreparedQuery("SELECT booking.*, RecurringType FROM booking, booking_info WHERE booking.ParentID = booking_info.ParentID AND booking.RoomID = ? AND booking.Begin > (SELECT DATE_SUB(CURRENT_DATE,INTERVAL 1 WEEK) ) ORDER BY booking.Begin", roomID)
	defer rows.Close()

	if err != nil {
		return []models.BookingModel{}, err
	}

	var bookings []models.BookingModel
	var booking models.BookingModel
	for rows.Next() {
		var tmpBegin, tmpEnd string
		rows.Scan(&booking.BookingID, &booking.ParentID, &booking.Title, &booking.RoomID, &booking.UserID, &tmpBegin, &tmpEnd, &booking.RecurringType, )

		booking.Begin = nil
		booking.End = nil
		booking.Begin = append(booking.Begin, tmpBegin)
		booking.End = append(booking.End, tmpEnd)

		bookings = append(bookings, booking)
	}
	return bookings, err
}

func (repositorys *BookingRepository) GetBookingByID(bookingID string) (models.BookingModel, error) {
	// TODO SQL Query anpassen
	rows, err := repositorys.DBHandler.PreparedQuery("SELECT booking.*, RecurringType FROM booking, booking_info WHERE booking.ParentID = booking_info.ParentID AND booking.BookingID = ?", bookingID)
	defer rows.Close()

	if err != nil {
		return models.BookingModel{}, err
	}

	var booking models.BookingModel
	rows.Next()
	var tmpBegin, tmpEnd string

	rows.Scan(&booking.BookingID, &booking.ParentID, &booking.Title, &booking.RoomID, &booking.UserID, &tmpBegin, &tmpEnd, &booking.RecurringType)

	booking.Begin = nil
	booking.End = nil
	booking.Begin = append(booking.Begin, tmpBegin)
	booking.End = append(booking.End, tmpEnd)

	return booking, err
}

func (repositorys *BookingRepository) CreateBooking(booking models.BookingModel) ([]models.BookingModel, error) {
	// FIXME hier ist das Problem, dass Query(), *Rows zurückliefert worunter man die lastinserid() nicht aufrufen kann
	// abhilfe wäre hier Exec() zu verwenden, dort wird ein *Stmt zurückgeliefert.

location, err := time.LoadLocation("Europe/Berlin")

if err != nil {
	fmt.Println(err)
}

beginTime , errParse:= time.ParseInLocation("2006-01-02 15:04:05",booking.Begin[0], location)

if errParse != nil {
	fmt.Println(err)
}

if (!beginTime.Add(time.Minute * 2).Before(time.Now())){

	var bookings []models.BookingModel
	tmpBooking := booking
	fmt.Println(booking)
	result, err := repositorys.DBHandler.PreparedQueryResult("INSERT INTO booking_info (RecurringType) VALUES(?)", booking.RecurringType)
	if err != nil {
		return []models.BookingModel{}, err
	}
	lastInsertBookingInfoID, _ := result.LastInsertId()


	for index, _ := range booking.Begin {
		result, err := repositorys.DBHandler.PreparedQueryResult("INSERT INTO booking (ParentID, Title, RoomID, UserID, Begin, End) VALUES(?, ?, ?, ?, ?, ?)",

		lastInsertBookingInfoID, booking.Title, booking.RoomID, booking.UserID, booking.Begin[index], booking.End[index])
		if err != nil {
			return []models.BookingModel{}, err
		}

		lastInsertID, _ := result.LastInsertId()

		//Convert the last inserted ID from the Databse which is int64 to string for the struct
		s := strconv.FormatInt(lastInsertID, 10)
		tmpBooking.BookingID = s
		tmpBooking.Begin = []string{booking.Begin[index]}
		tmpBooking.End = []string{booking.End[index]}

		bookings = append(bookings, tmpBooking)
	}

	return bookings, nil
} else {
	var bookings []models.BookingModel
	bookings = append(bookings, booking)
	err := errors.New("Can not create booking in the past")
	return bookings, err
}

}

func (repositorys *BookingRepository) UpdateBooking(bookingID string, bookingToUpdate models.BookingModel) (models.BookingModel, error) {
	// TODO ALLE REPOS mit der update funktion anpassen -> SQL QUERY

	location, err := time.LoadLocation("Europe/Berlin")

if err != nil {
	fmt.Println(err)
}

	beginTime ,_:= time.ParseInLocation("2006-01-02 15:04:00",bookingToUpdate.Begin[0], location)
	if (!beginTime.Before(time.Now())){

	rows, err := repositorys.DBHandler.PreparedQuery("UPDATE booking set Title= ?, RoomID= ?, Begin= ?, End= ? where BookingID= ?;",
		bookingToUpdate.Title,	
		bookingToUpdate.RoomID,
		bookingToUpdate.Begin[0],
		bookingToUpdate.End[0],
		bookingID)
	defer rows.Close()

	if err != nil {
		fmt.Println(err)
		return models.BookingModel{}, err
	}

	var booking models.BookingModel
	rows.Next()
	rows.Scan(&booking.BookingID, &booking.ParentID, &booking.Title, &booking.RoomID, &booking.UserID, &booking.Begin, &booking.End, &booking.RecurringType)

	return booking, err
}
	err = errors.New("Can not update booking to the past")
	return bookingToUpdate, err
}

func (repositorys *BookingRepository) DeleteBooking(bookingID string) (models.BookingModel, error) {
	var booking models.BookingModel

	countRow, _ := repositorys.DBHandler.PreparedQuery("Select count(*) FROM booking WHERE ParentID = (SELECT ParentID FROM booking WHERE BookingID =?)", bookingID)
	defer countRow.Close()
	var count int
	countRow.Next()
	countRow.Scan(&count)
	if (count > 1){

	fmt.Println(count)

	rows, err := repositorys.DBHandler.PreparedQuery("DELETE FROM booking WHERE BookingID = ?", bookingID)
	defer rows.Close()

	if err != nil {
		return models.BookingModel{}, err
	}

	rows.Next()
	// FIXME should i return here the data object?
	rows.Scan(&booking.BookingID, &booking.ParentID, &booking.Title, &booking.RoomID, &booking.UserID, &booking.Begin, &booking.End, &booking.RecurringType)
	
	return booking, err
	}	else {
		rows, err := repositorys.DBHandler.PreparedQuery("DELETE FROM booking_info WHERE ParentID = (SELECT ParentID FROM booking WHERE BookingID =?)", bookingID)
		defer rows.Close()
		if err != nil {
			return models.BookingModel{}, err
		}
	
		rows.Next()
		// FIXME should i return here the data object?
		rows.Scan(&booking.BookingID, &booking.ParentID, &booking.Title, &booking.RoomID, &booking.UserID, &booking.Begin, &booking.End, &booking.RecurringType)
		return booking, err

	}

}

func (repositorys *BookingRepository) DeleteRecurringBooking(parentID string) (models.BookingModel, error) {


	rows, err := repositorys.DBHandler.PreparedQuery("DELETE FROM booking_info WHERE ParentID = ?", parentID)
	defer rows.Close()

	if err != nil {

		return models.BookingModel{}, err
	}

	var booking models.BookingModel
	rows.Next()
	// FIXME should i return here the data object?
	rows.Scan(&booking.BookingID, &booking.Title, &booking.RoomID, &booking.UserID, &booking.Begin, &booking.End, &booking.RecurringType)

	return booking, err
}