package ws

import (
	"github.com/gorilla/websocket"
	"time"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"fmt"
)

const (
	writeWait  = 5 * time.Second
	pingPeriod = 5 * time.Second
)

type Client struct {
	Id   string
	Conn *websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	params := mux.Vars(r)

	client := &Client{
		Id:   params["id"],
		Conn: conn,
	}
	H.register <- client
	go client.listenHub()
	go client.isConnected()
}

func (c *Client) listenHub() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		H.unregister <- c
		c.Conn.Close()
	}()
	for {
		select {
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) isConnected() {
	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			return
		}
		if tv, ok := H.Clients["tv"]; ok {
			fmt.Println("TV presents: ", tv)
			tv.Conn.WriteMessage(websocket.BinaryMessage, msg)
		}
		//fmt.Println(string(msg))
		//fmt.Println(msg)
	}
}
