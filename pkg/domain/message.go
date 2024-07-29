package domain

type Message struct {
	ID         string `json:"id" bson:"_id,omitempty"`
	SenderID   string `json:"sender_id"`
	ReceiverID string `json:"receiver_id"`
	Content    string `json:"content"`
	Timestamp  int64  `json:"timestamp"`
}
