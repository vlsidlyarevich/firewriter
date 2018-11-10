package app

import (
	"github.com/gorilla/websocket"
	"github.com/vlsidlyarevich/firewriter/internal/app/firewriter/model"
	"log"
	"net/http"
)

type Server struct {
	clients   map[*websocket.Conn]bool
	broadcast chan model.Message
	upgrader  websocket.Upgrader
}

func (server *Server) Init() {
	server.clients = make(map[*websocket.Conn]bool)
	server.broadcast = make(chan model.Message)

	// Configure the upgrader
	server.upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}

func (server *Server) Start() {
	// Create a simple file server
	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/", fs)

	// Configure websocket route
	//http.HandleFunc("/ws", handleConnections)
	//
	//go handleMessages()

	// Start the server on localhost port 8000 and log any errors
	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
