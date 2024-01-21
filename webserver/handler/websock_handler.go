package handler

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/http"
	"simpleboard2/resource"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		println("WebSocket Updrade Failed:", err)
		return
	}
	defer conn.Close()

	systemResource := resource.NewSystemResource()

	for {
		systemResource.GetSystemResource()
		jsonString, err := json.MarshalIndent(systemResource, "", "  ")
		if err != nil {
			println("json marshal error:", err)
			break
		}
		err = conn.WriteMessage(websocket.TextMessage, jsonString)
		if err != nil {
			println("Write Message Error:", err)
			break
		}
	}

}
