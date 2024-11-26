package memory

import (
	"errors"
	"reddit-clone/internal/models"
	"sync"
)

//goland:noinspection SpellCheckingInspection
type MemoryStore struct {
	users      map[string]*models.User
	subreddits map[string]*models.Subreddit
	posts      map[string]*models.Post
	votes      map[string]map[string]bool // targetID -> userID -> upvote/downvote
	mu         sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		users:      make(map[string]*models.User),
		subreddits: make(map[string]*models.Subreddit),
		posts:      make(map[string]*models.Post),
		votes:      make(map[string]map[string]bool),
	}
}

// User operations
func (m *MemoryStore) CreateUser(user *models.User) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.users[user.ID]; exists {
		return errors.New("user already exists")
	}

	m.users[user.ID] = user
	return nil
}

func (m *MemoryStore) GetUser(id string) (*models.User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	user, exists := m.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// Subreddit operations
func (m *MemoryStore) CreateSubreddit(subreddit *models.Subreddit) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.subreddits[subreddit.ID]; exists {
		return errors.New("subreddit already exists")
	}

	m.subreddits[subreddit.ID] = subreddit
	return nil
}

func (m *MemoryStore) GetSubreddit(id string) (*models.Subreddit, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	subreddit, exists := m.subreddits[id]
	if !exists {
		return nil, errors.New("subreddit not found")
	}
	return subreddit, nil
}

func (m *MemoryStore) JoinSubreddit(subredditID, userID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	subreddit, exists := m.subreddits[subredditID]
	if !exists {
		return errors.New("subreddit not found")
	}

	if subreddit.Members == nil {
		subreddit.Members = make(map[string]bool)
	}
	subreddit.Members[userID] = true
	return nil
}

func (m *MemoryStore) LeaveSubreddit(subredditID, userID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	subreddit, exists := m.subreddits[subredditID]
	if !exists {
		return errors.New("subreddit not found")
	}

	delete(subreddit.Members, userID)
	return nil
}

// Post operations
func (m *MemoryStore) CreatePost(post *models.Post) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.posts[post.ID]; exists {
		return errors.New("post already exists")
	}

	m.posts[post.ID] = post
	return nil
}

func (m *MemoryStore) GetPost(id string) (*models.Post, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	post, exists := m.posts[id]
	if !exists {
		return nil, errors.New("post not found")
	}
	return post, nil
}

func (m *MemoryStore) GetSubredditPosts(subredditID string) ([]*models.Post, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var posts []*models.Post
	for _, post := range m.posts {
		if post.SubredditID == subredditID {
			posts = append(posts, post)
		}
	}
	return posts, nil
}

// Vote operations
func (m *MemoryStore) Vote(targetID, userID string, isUpvote bool) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.votes[targetID]; !exists {
		m.votes[targetID] = make(map[string]bool)
	}

	m.votes[targetID][userID] = isUpvote

	// Update karma for the target (post or comment)
	if post, exists := m.posts[targetID]; exists {
		if isUpvote {
			post.Karma++
		} else {
			post.Karma--
		}
	}

	return nil
}
