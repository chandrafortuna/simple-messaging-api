package main

import (
	"log"
	"net/http"

	m "github.com/chandrafortuna/simple-messaging-api/domain/message"
	h "github.com/chandrafortuna/simple-messaging-api/handler"
	"github.com/gorilla/mux"
)

func main() {
	//register service
	messageRepository := m.NewRepository([]*m.Message{})
	messageService := m.NewService(messageRepository)
	handler := h.NewHandler(messageService)

	router := mux.NewRouter()
	router.HandleFunc("/chat", handler.Send).Methods("POST")
	router.HandleFunc("/chat", handler.GetAll).Methods("GET")
	router.HandleFunc("/ws", handler.WebsocketEndpoint)
	log.Fatal(http.ListenAndServe(":8000", router))
}
