package application

import (
	"encoding/json"
	"log"
	"api-producer/src/message/Domain"
)

type MessageService struct {
	Producer Domain.Producer
}

func NewMessageService(producer Domain.Producer) *MessageService {
	return &MessageService{Producer: producer}
}

func (s *MessageService) HandleMessage(message Domain.Message) error {

	msgBytes, err := json.Marshal(message)
	if err != nil {
		log.Printf("Mensaje de error al ordenar: %s", err)
		return err
	}

	err = s.Producer.SendMessageToQueue("machine_status", msgBytes)
	if err != nil {
		log.Printf("Error al enviar el mensaje a RabbitMQ: %s", err)
		return err
	}

	log.Printf("Mensaje enviado exitosamente a RabbitMQ: %s", message.Content)
	return nil
}