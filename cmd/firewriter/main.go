package main

import "github.com/vlsidlyarevich/firewriter/web/app"

func main() {
	var server = app.Server{}
	server.Init()
	server.Start()
}

//
//func handleConnections(w http.ResponseWriter, r *http.Request) {
//	// Upgrade initial GET request to a websocket
//	ws, err := upgrader.Upgrade(w, r, nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	// Make sure we close the connection when the function returns
//	defer ws.Close()
//	clients[ws] = true
//
//	for {
//		var msg model.Message
//		// Read in a new message as JSON and map it to a Message object
//		err := ws.ReadJSON(&msg)
//		if err != nil {
//			log.Printf("error: %v", err)
//			delete(clients, ws)
//			break
//		}
//		// Send the newly received message to the broadcast channel
//		broadcast <- msg
//	}
//}
//
//func handleMessages() {
//	for {
//		// Grab the next message from the broadcast channel
//		msg := <-broadcast
//		// Send it out to every client that is currently connected
//		for client := range clients {
//			err := client.WriteJSON(msg)
//			if err != nil {
//				log.Printf("error: %v", err)
//				client.Close()
//				delete(clients, client)
//			}
//		}
//	}
//}
