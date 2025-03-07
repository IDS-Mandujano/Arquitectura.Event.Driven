package Domain

type Producer interface {
	SendMessageToQueue(queueName string, message []byte) error
}
