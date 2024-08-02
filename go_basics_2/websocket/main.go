package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	indexHtml, _ := os.Open("index.html")
	defer indexHtml.Close()
	indexData, _ := io.ReadAll(indexHtml)
	fmt.Fprint(w, string(indexData))
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, _ := upgrader.Upgrade(w, r, nil)
	defer conn.Close()
	mType, bts, _ := conn.ReadMessage()
	fmt.Println(string(bts))
	conn.WriteMessage(mType, []byte("message from backend"))
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/ws", WebSocketHandler)
	http.ListenAndServe(":8888", nil)
}
