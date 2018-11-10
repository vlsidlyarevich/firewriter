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
	server.upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}

func (server *Server) InitHandlers() {
	handleConnections := func(w http.ResponseWriter, r *http.Request) {
		// Upgrade initial GET request to a websocket
		ws, err := server.upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}
		// Make sure we close the connection when the function returns
		defer ws.Close()
		server.clients[ws] = true

		for {
			var msg model.Message
			// Read in a new message as JSON and map it to a Message object
			err := ws.ReadJSON(&msg)
			if err != nil {
				log.Printf("error: %v", err)
				delete(server.clients, ws)
				break
			}
			// Send the newly received message to the broadcast channel
			server.broadcast <- msg
		}
	}

	//Configure websocket route
	http.HandleFunc("/ws", handleConnections)
}

func (server *Server) Start(staticDir, port string) {
	// Create a simple file server
	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/", fs)

	go handleMessages(server)

	// Start the server on localhost port 8000 and log any errors
	log.Println("http server started on :" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleMessages(server *Server) {
	for {
		// Grab the next message from the broadcast channel
		msg := <-server.broadcast
		// Send it out to every client that is currently connected
		for client := range server.clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(server.clients, client)
			}
		}
	}
}
