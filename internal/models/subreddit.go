package models

type Subreddit struct {
	ID          string
	Name        string
	Description string
	CreatorID   string
	Members     map[string]bool // user_id -> membership status
	Created     int64
}
