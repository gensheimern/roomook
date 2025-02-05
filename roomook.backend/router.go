package main

import (
	"fmt"
	"net/http"
	"log"
	"sync"
	"bytes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"jacob.de/roomook/roomook.backend/controllers"
	"time"

)

//IGinRouter Interface
type IGinRouter interface {
	SetupRouter() *gin.Engine
}

type router struct{}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Conn is an middleman between the websocket connection and the hub.
type Conn struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

func (router *router) SetupRouter() *gin.Engine {
	roomController := ServiceContainer().InjectRoomController()
	bookingController := ServiceContainer().InjectBookingController()

	r := gin.Default()
	r.Use(cors.Default())

	user := r.Group("/user/")
	{
		user.GET("/loggedinuser/", controllers.GetLoggedInUser)

	}



	room := r.Group("/room/")
	{
		room.GET("/", roomController.GetAllRoom)
		room.GET("/:id", roomController.GetRoomByID)
		room.POST("/", roomController.CreateRoom)
		room.PUT("/:id", roomController.UpdateRoom)
		room.DELETE("/:id", roomController.DeleteRoom)

	}

	booking := r.Group("/booking/")
	{
		booking.GET("/", bookingController.GetAllBooking)
		booking.GET("/:id", bookingController.GetBookingByID)
		booking.POST("/", bookingController.CreateBooking)
		booking.PUT("/:id", bookingController.UpdateBooking)
		booking.DELETE("/:id", bookingController.DeleteBooking)
		booking.DELETE("/:id/recurrent/", bookingController.DeleteRecurringBooking)

	}

	bookingByRoom := r.Group("/b/r/")
	{
		bookingByRoom.GET("/:id", bookingController.GetAllBookingByRoomID)
	}

	roomByDate := r.Group("/r/d")
	{
		roomByDate.POST("/", roomController.GetRoomByDate)
	}

	websocket := r.Group("/ws/")
	{
		websocket.GET("/", func(ctx *gin.Context) {
			serveWs( ctx.Writer, ctx.Request)
		})
	}

	return r
}
/*
func wshandler(ctx *gin.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Println("websooooooooooooooockets")

	wsupgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	clients[conn] = true

	for {
		var wsm WsMessage
		 err := conn.ReadJSON(&wsm)
		if err != nil {
			fmt.Println(err)
			delete(clients, conn)
			break
		}

		broadcast <- wsm
	}

}

func handleMessages() {

	for {

		wsm := <-broadcast
		for client := range clients {
			fmt.Println(client)

			var i int
			err := client.WriteJSON(wsm)
			i = i + 1
			if err != nil {
				fmt.Println(err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
*/


// readPump pumps messages from the websocket connection to the hub.
func (c *Conn) readPump() {
	defer func() {
		hub.unregister <- c
		c.ws.Close()
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				fmt.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		hub.broadcast <- message
	}
}

// write writes a message with the given message type and payload.
func (c *Conn) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}

// writePump pumps messages from the hub to the websocket connection.
func (c *Conn) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				// The hub closed the channel.
				c.write(websocket.CloseMessage, []byte{})
				return
			}

			c.ws.SetWriteDeadline(time.Now().Add(writeWait))
			w, err := c.ws.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	conn := &Conn{send: make(chan []byte, 256), ws: ws}
	hub.register <- conn
	go conn.writePump()
	conn.readPump()
}

var (
	m          *router
	routerOnce sync.Once
)

func GinRouter() IGinRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}
