package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anurag4way/go-concurrency/types"
	"github.com/gorilla/websocket"
)

func main() {
	recv := NewDataReceiver()
	http.HandleFunc("/ws", recv.handleWS)
	http.ListenAndServe(":30000", nil)
	fmt.Println("data receiver working fine")
}

type DataReceiver struct {
	msgch chan types.OBUDATA
	conn  *websocket.Conn
}

func NewDataReceiver() *DataReceiver {
	return &DataReceiver{
		msgch: make(chan types.OBUDATA, 128),
	}
}

func (dr *DataReceiver) handleWS(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{
		ReadBufferSize:  1028,
		WriteBufferSize: 1028,
	}

	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	dr.conn = conn
	go dr.wsRecieveLoop()
}
func (dr *DataReceiver) wsRecieveLoop() {
	fmt.Println("new obu client connected")
	for {
		var data types.OBUDATA
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println("read error:", err)
			continue
		}
		fmt.Printf("received OBU data from [%d]:: <lat %.2f, Long %.2f>\n", data.OBUID, data.Lat, data.Long)
	}
}
