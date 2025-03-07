package application

import (
	"encoding/json"
	"fmt"
	"net/http"
	"api-producer/src/message/Domain"
)

type MessageHandler struct {
	MessageService *MessageService
}

func NewMessageHandler(messageService *MessageService) *MessageHandler {
	return &MessageHandler{MessageService: messageService}
}

func (h *MessageHandler) HandleMessage(w http.ResponseWriter, r *http.Request) {
	var msg Domain.Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding message: %s", err), http.StatusBadRequest)
		return
	}

	err := h.MessageService.HandleMessage(msg)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error processing message: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message successfully sent to RabbitMQ"))
}