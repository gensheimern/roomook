package main

import (
	"database/sql"
	"sync"
	_ "github.com/go-sql-driver/mysql"
	
	"jacob.de/roomook/roomook.backend/controllers"
	"jacob.de/roomook/roomook.backend/infrastructures"
	"jacob.de/roomook/roomook.backend/repositories"
	"jacob.de/roomook/roomook.backend/services"
	"jacob.de/roomook/roomook.backend/config"
)


type IServiceContainer interface {
	InjectRoomController() controllers.RoomController
	InjectBookingController() controllers.BookingController
}

type kernel struct{}

func (k *kernel) InjectRoomController() controllers.RoomController {

	cfg  := config.ReadConfig()

	sqlConn, _ := sql.Open(cfg.Db, cfg.Db_user + ":" + cfg.Db_password + "@tcp(" + cfg.Db_host + ")/" + cfg.Db_name)
	mysqlHandler := &infrastructures.MysqlHandler{}
	mysqlHandler.Conn = sqlConn

	roomRepository := &repositories.RoomRepository{mysqlHandler}
	roomService := &services.RoomService{roomRepository}
	roomController := controllers.RoomController{roomService}
	return roomController
}

func (k *kernel) InjectBookingController() controllers.BookingController {

	cfg  := config.ReadConfig()

	sqlConn, _ := sql.Open(cfg.Db, cfg.Db_user + ":" + cfg.Db_password + "@tcp(" + cfg.Db_host + ")/" + cfg.Db_name)
	mysqlHandler := &infrastructures.MysqlHandler{}
	mysqlHandler.Conn = sqlConn

	bookingRepository := &repositories.BookingRepository{mysqlHandler}
	bookingService := &services.BookingService{bookingRepository}
	bookingController := controllers.BookingController{bookingService}
	return bookingController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
