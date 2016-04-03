package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"yelpmain"
)

var connections map[*websocket.Conn]bool

func main() {
	port := flag.Int("port", 8000, "")
	dir := flag.String("indexLocation", "", "")

	flag.Parse()
	connections = make(map[*websocket.Conn]bool)

	fs := http.Dir(*dir)
	fileHandler := http.FileServer(fs)
	http.Handle("/", fileHandler)

	http.HandleFunc("/sock", handler)
	log.Printf("FTHacks server started on port %d\n", *port)

	addr := fmt.Sprintf("localhost:%d", *port)
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	_, ok := err.(websocket.HandshakeError)
	if ok {
		http.Error(w, "error", 400)
		return
	} else if err != nil {
		log.Println(err)
		return
	}

	connections[conn] = true

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			delete(connections, conn)
			conn.Close()
			return
		}
		s := strings.Split(string(msg), ", ")
		init(s[0], s[1])
		sendAll(msg)
	}
}

func sendAll(message []byte) {
	for conn := range connections {
		err := conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			delete(connections, conn)
			conn.Close()
		}
	}
}
