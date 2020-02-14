package handler

import (
	"log"
	"net/http"

	m "github.com/chandrafortuna/simple-messaging-api/domain/message"
	"github.com/gorilla/websocket"
)

type Handler struct {
	service m.Service
	socket  *websocket.Conn
}

func NewHandler(s m.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) Send(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("message")
	res, err := h.service.Send(msg)
	if err != nil {
		// respondErr(w, r, http.StatusInternalServerError, "failed to write message: ", err)
		ReplyError(&w, http.StatusInternalServerError, "failed to write message: ", err)
		return
	}

	if h.socket != nil {
		if err := h.socket.WriteMessage(websocket.TextMessage, []byte(res.Text)); err != nil {
			log.Println("Write message failed:", err)
			return
		}
	}

	log.Println(res.Text)

	// respond(w, r, http.StatusCreated, &res)
	ReplySuccess(&w, &res)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := h.service.GetAll()
	if err != nil {
		// respondErr(w, r, http.StatusInternalServerError, "failed to get all message: ", err)
		ReplyError(&w, http.StatusInternalServerError, "failed to get all message: ", err)
		return
	}

	// respond(w, r, http.StatusOK, &res)
	ReplySuccess(&w, res)
}

func (h *Handler) wsReader() {
	for {
		messageType, p, err := h.socket.ReadMessage()
		if err != nil {
			log.Println("Read Message failed:", err)
			return
		}

		log.Println(string(p))

		if err := h.socket.WriteMessage(messageType, p); err != nil {
			log.Println("Write message failed:", err)
			return
		}
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *Handler) WebsocketEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrader Err:", err)
		return
	}
	log.Println("websocket connection success.")
	defer ws.Close()
	h.socket = ws
	h.wsReader()
}
