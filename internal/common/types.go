package common

import (
	"math/rand"
	"time"
)

// Messages
type SimulateAction struct {
	Timestamp time.Time
	Subreddit string
}

type ConnectionStatus struct {
	Connected bool
}

// Behavior types
type ClientBehavior struct {
	PostProbability    float64
	CommentProbability float64
	VoteProbability    float64
	JoinProbability    float64
	ActiveHours        []int
	Persona            string
}

// Distribution types
type SimulationDistribution struct {
	Subreddits  []string
	PostWeights map[string]float64
}

// NewSimulationDistribution creates a new distribution with initial weights
func NewSimulationDistribution(subreddits []string) *SimulationDistribution {
	weights := make(map[string]float64)
	for _, subreddit := range subreddits {
		// Initialize with random weights following Zipf-like distribution
		weights[subreddit] = 1.0 / float64(1+rand.Intn(100))
	}

	return &SimulationDistribution{
		Subreddits:  subreddits,
		PostWeights: weights,
	}
}

// Helper functions
func (d *SimulationDistribution) GetRandomSubreddit() string {
	if len(d.Subreddits) == 0 {
		return ""
	}
	return d.Subreddits[rand.Intn(len(d.Subreddits))]
}

func (d *SimulationDistribution) ShouldCreatePost(subreddit string) bool {
	if weight, exists := d.PostWeights[subreddit]; exists {
		return rand.Float64() < weight
	}
	return false
}

func (d *SimulationDistribution) GetSubreddits() []string {
	return d.Subreddits
}

func (d *SimulationDistribution) UpdateWeight(subreddit string, delta float64) {
	if weight, exists := d.PostWeights[subreddit]; exists {
		d.PostWeights[subreddit] = weight + delta
		if d.PostWeights[subreddit] > 1.0 {
			d.PostWeights[subreddit] = 1.0
		}
	}
}

// PingMessage is used to verify connection
type PingMessage struct{}

// PongMessage is the response to a ping
type PongMessage struct{}

func GenerateUserBehavior() *ClientBehavior {
	personas := []string{"Lurker", "Casual", "PowerUser"}
	persona := personas[rand.Intn(len(personas))]

	behavior := &ClientBehavior{
		Persona:     persona,
		ActiveHours: generateActiveHours(),
	}

	switch persona {
	case "Lurker":
		behavior.PostProbability = 0.05
		behavior.CommentProbability = 0.1
		behavior.VoteProbability = 0.85
		behavior.JoinProbability = 0.05
	case "Casual":
		behavior.PostProbability = 0.2
		behavior.CommentProbability = 0.3
		behavior.VoteProbability = 0.5
		behavior.JoinProbability = 0.2
	case "PowerUser":
		behavior.PostProbability = 0.4
		behavior.CommentProbability = 0.4
		behavior.VoteProbability = 0.2
		behavior.JoinProbability = 0.4
	}

	return behavior
}

func generateActiveHours() []int {
	numHours := 8 + rand.Intn(8) // 8-16 active hours
	hours := make([]int, numHours)

	// Generate random active hours (0-23)
	for i := 0; i < numHours; i++ {
		hours[i] = rand.Intn(24)
	}

	return hours
}

type ActionType string

const (
	PostAction    ActionType = "post"
	CommentAction ActionType = "comment"
	VoteAction    ActionType = "vote"
	JoinAction    ActionType = "join"
)

type Action struct {
	Type      ActionType
	UserID    string
	Content   string
	Timestamp time.Time
}

type ActionResponse struct {
	Success     bool
	ActiveUsers int
	Error       error
}
