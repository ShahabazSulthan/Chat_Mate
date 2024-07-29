package repositories

import (
	"chatting/pkg/domain"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageRepository struct {
	Collection *mongo.Collection
}

func (m *MessageRepository) Create(message *domain.Message) error {
	_, err := m.Collection.InsertOne(context.Background(), message)
	return err
}

func (m *MessageRepository) FindBySenderAndReceiver(senderID, receiverID string) ([]*domain.Message, error) {
	var messages []*domain.Message
	filter := bson.M{
		"$or": []bson.M{
			{"senderid": senderID, "receiverid": receiverID},
			{"senderid": receiverID, "receiverid": senderID},
		},
	}
	fmt.Println("Hhhhh", senderID, receiverID)
	cursor, err := m.Collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		var message domain.Message
		if err := cursor.Decode(&message); err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}
	return messages, nil
}
