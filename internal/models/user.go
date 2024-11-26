package models

type User struct {
	ID       string
	Username string
	Password string
	Karma    int32
	Created  int64
}
