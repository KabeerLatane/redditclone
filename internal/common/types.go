package common

import (
	"math/rand"
	"time"
)

// Messages
type SimulateAction struct {
	Timestamp time.Time
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
