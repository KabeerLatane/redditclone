// internal/simulation/
package simulation

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type SimulationDistribution struct {
	zipf        *rand.Zipf
	subreddits  []string
	postWeights map[string]float64
}

func NewSimulationDistribution(numSubreddits int, alpha float64) *SimulationDistribution {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Initialize Zipf distribution
	zipf := rand.NewZipf(random, alpha, 1.0, uint64(numSubreddits))

	// Generate subreddits
	subreddits := make([]string, numSubreddits)
	postWeights := make(map[string]float64)

	for i := 0; i < numSubreddits; i++ {
		subreddits[i] = generateSubredditName(i)
		// Larger subreddits get more posts (following Zipf)
		weight := 1.0 / math.Pow(float64(i+1), alpha)
		postWeights[subreddits[i]] = weight
	}

	return &SimulationDistribution{
		zipf:        zipf,
		subreddits:  subreddits,
		postWeights: postWeights,
	}
}

func (sd *SimulationDistribution) GetRandomSubreddit() string {
	index := sd.zipf.Uint64()
	if int(index) >= len(sd.subreddits) {
		index = uint64(len(sd.subreddits) - 1)
	}
	return sd.subreddits[index]
}

func (sd *SimulationDistribution) ShouldCreatePost(subreddit string) bool {
	weight := sd.postWeights[subreddit]
	return rand.Float64() < weight
}

func generateSubredditName(index int) string {
	topics := []string{"gaming", "tech", "science", "music", "movies", "books", "sports", "news", "food", "art"}
	subtopics := []string{"discussion", "meta", "help", "pics", "videos", "memes"}

	topic := topics[index%len(topics)]
	subtopic := subtopics[(index/len(topics))%len(subtopics)]

	return fmt.Sprintf("r_%s_%s_%d", topic, subtopic, index)
}
