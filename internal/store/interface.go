// store/interface.go
package store

import "reddit-clone/internal/models"

type Store interface {
	// User operations
	CreateUser(user *models.User) error
	GetUser(id string) (*models.User, error)

	// Subreddit operations
	CreateSubreddit(subreddit *models.Subreddit) error
	GetSubreddit(id string) (*models.Subreddit, error)
	JoinSubreddit(subredditID, userID string) error
	LeaveSubreddit(subredditID, userID string) error

	// Post operations
	CreatePost(post *models.Post) error
	GetPost(id string) (*models.Post, error)
	GetSubredditPosts(subredditID string) ([]*models.Post, error)

	// Comment operations
	AddComment(comment *models.Comment) error
	GetComments(postID string) ([]*models.Comment, error)

	// Message operations
	SendMessage(message *models.DirectMessage) error
	GetMessages(userID string) ([]*models.DirectMessage, error)

	// Vote operations
	Vote(targetID, userID string, isUpvote bool) error
}
