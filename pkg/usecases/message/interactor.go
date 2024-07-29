// /internal/usecases/message/interactor.go
package message

import (
	"chatting/pkg/domain"
	"chatting/pkg/repositories"
)

type MessageInteractor struct {
	MessageRepository repositories.MessageRepository
}

func (m *MessageInteractor) SendMessage(message *domain.Message) error {
	return m.MessageRepository.Create(message)
}

func (m *MessageInteractor) GetMessages(senderID, receiverID string) ([]*domain.Message, error) {
	return m.MessageRepository.FindBySenderAndReceiver(senderID, receiverID)
}
