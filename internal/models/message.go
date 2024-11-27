package models

type DirectMessage struct {
	ID        string
	FromID    string
	ToID      string
	Content   string
	Timestamp int64
}
