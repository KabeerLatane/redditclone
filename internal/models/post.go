package models

type Post struct {
	ID          string
	SubredditID string
	AuthorID    string
	Title       string
	Content     string
	Karma       int32
	Created     int64
	Votes       map[string]bool // user_id -> upvote(true)/downvote(false)
}
