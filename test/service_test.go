package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	m "github.com/chandrafortuna/simple-messaging-api/domain/message"
	h "github.com/chandrafortuna/simple-messaging-api/handler"
)

var (
	messageRepository = m.NewRepository([]*m.Message{})
	messageService    = m.NewService(messageRepository)
	messageHandler    = h.NewHandler(messageService)
	contentMessage    = "This is a test."
	expectedMessage   = &m.Message{
		Text: contentMessage,
	}
)

func TestSendMessage(t *testing.T) {
	req, err := http.NewRequest("POST", "/chat?message="+contentMessage, nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(messageHandler.Send)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"text":"This is a test."}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetAll(t *testing.T) {
	req, err := http.NewRequest("GET", "/chat", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(messageHandler.GetAll)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
