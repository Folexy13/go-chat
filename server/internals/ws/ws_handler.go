package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct{
	hub *Hub
}

func NewHandler(h *Hub) *Handler{
	return &Handler{
		hub:h,
	}
}

type CreateRoomReq struct{
	ID string `json:"id"`
	Name string `json:"name"`
}
func (h *Handler) CreateRoom(c *gin.Context){
var req CreateRoomReq
if err := c.ShouldBindJSON(&req);err!=nil{
	c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	return
}
h.hub.Rooms[req.ID]=&Room{
	ID:req.ID,
	Name:req.Name,
	Clients: make(map[string]*Client),
}
c.JSON(http.StatusOK,req)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool{
		return true
	},
}

func (h *Handler) JoinRoom(c *gin.Context){
conn,err:=upgrader.Upgrade(c.Writer, c.Request,nil)
if err !=nil{
	c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	return
}
roomId:=c.Param("roomId")
clientId:= c.Query("userId")
username := c.Query("username")
cl:= &Client{
	Conn: conn,
	Message: make(chan *Message,10),
	ID:clientId,
	RoomID: roomId,
	Username: username,
}
m:= &Message{
	Content: "A new user has joined the room",
	RoomID:roomId,
	Username: username,
}

//Register a new client through the register channel
h.hub.Register <- cl

//Broadast message
h.hub.Broadcast <- m
}