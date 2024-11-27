package models

type Comment struct {
	ID       string
	PostID   string
	ParentID string // empty if top-level comment
	AuthorID string
	Content  string
	Created  int64
	Children []string // IDs of child comments
}
