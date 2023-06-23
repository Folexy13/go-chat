package ws

type Room struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct{
	Rooms map[string]*Room
	Register chan *Client
	Unregister chan *Client
	Broadcast chan *Message
}

func NewHub() *Hub{
	return &Hub{
		Rooms:make(map[string]*Room),
	}
}