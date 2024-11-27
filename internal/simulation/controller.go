package simulation

import (
	"fmt"
	protoactor "github.com/asynkron/protoactor-go/actor"
	"log"
	"math"
	"math/rand"
	"reddit-clone/internal/actor"
	"reddit-clone/internal/common"
	"reddit-clone/pkg/metrics"
	"reddit-clone/pkg/utils"
	"sync"
	"sync/atomic"
	"time"
)

var mutex = &sync.Mutex{}

type SimulationController struct {
	system           *protoactor.ActorSystem
	enginePID        *protoactor.PID
	clients          []*protoactor.PID
	zipf             *rand.Zipf
	metrics          *metrics.RedditMetrics
	distribution     *common.SimulationDistribution
	subredditWeights map[string]float64
	postPopularity   map[string]int
	mutex            sync.Mutex
	baseCount        atomic.Int32
}

func NewSimulationController(system *protoactor.ActorSystem, enginePID *protoactor.PID, metrics *metrics.RedditMetrics) *SimulationController {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	subreddits := []string{
		"technology", "gaming", "movies", "music", "books",
		"science", "sports", "news", "funny", "pics",
	}

	weights := make(map[string]float64)
	for i, subreddit := range subreddits {
		rank := float64(i + 1)
		// Improved Zipf calculation with better scaling
		weights[subreddit] = 1.0 / math.Pow(rank, 1.07)
	}

	return &SimulationController{
		system:           system,
		enginePID:        enginePID,
		clients:          make([]*protoactor.PID, 0),
		metrics:          metrics,
		zipf:             rand.NewZipf(r, 1.1, 1.0, 1000),
		distribution:     common.NewSimulationDistribution(subreddits),
		subredditWeights: weights,
		postPopularity:   make(map[string]int),
	}
}

func (s *SimulationController) Start(numClients int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.enginePID == nil {
		return fmt.Errorf("engine PID is nil")
	}
	// Get current base count
	currentCount := int(s.baseCount.Load())
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	s.zipf = rand.NewZipf(r, 1.1, 1.0, uint64(numClients))

	personaCounts := make(map[string]int)

	for i := 0; i < numClients; i++ {
		uniqueIndex := currentCount + i
		behavior := common.GenerateUserBehavior()
		personaCounts[behavior.Persona]++

		clientActor := actor.NewClientActor(
			utils.GenerateID(),
			utils.GenerateUsername(uniqueIndex),
			s.enginePID,
			behavior,
			s.metrics,
		)

		props := protoactor.PropsFromProducer(func() protoactor.Actor {
			return clientActor
		})

		pid, err := s.system.Root.SpawnNamed(props, fmt.Sprintf("client-%d", uniqueIndex))
		if err != nil {
			return fmt.Errorf("failed to spawn client actor: %v", err)
		}
		s.clients = append(s.clients, pid)
	}
	// Update base count atomically
	s.baseCount.Add(int32(numClients))
	for persona, count := range personaCounts {
		s.metrics.UpdatePersonaCount(persona, count)
	}

	go s.simulateActions()
	go s.simulateConnections()
	go s.reportMetrics()
	go s.updateWeights() // New goroutine for weight updates

	s.metrics.UpdateSimulationMetrics(float64(len(s.clients)), 1.0)

	return nil
}

func (s *SimulationController) simulateActions() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		clientIndex := s.zipf.Uint64()
		if int(clientIndex) >= len(s.clients) {
			continue
		}

		// Select subreddit using weighted random selection
		mutex.Lock()
		subreddit := s.getWeightedSubreddit()
		s.postPopularity[subreddit]++
		mutex.Unlock()

		s.system.Root.Send(s.clients[clientIndex], &common.SimulateAction{
			Timestamp: time.Now(),
			Subreddit: subreddit,
		})

	}
}

func (s *SimulationController) getWeightedSubreddit() string {
	total := 0.0
	for _, weight := range s.subredditWeights {
		total += weight
	}

	r := rand.Float64() * total
	cumulative := 0.0

	for subreddit, weight := range s.subredditWeights {
		cumulative += weight
		if r <= cumulative {
			return subreddit
		}
	}

	// Fallback to first subreddit
	for subreddit := range s.subredditWeights {
		return subreddit
	}
	return ""
}

func (s *SimulationController) updateWeights() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		mutex.Lock()
		// Adjust weights based on activity
		for subreddit := range s.subredditWeights {
			activity := float64(s.postPopularity[subreddit])
			if activity > 0 {
				s.subredditWeights[subreddit] *= (1.0 + (activity / 1000.0))
			}
		}

		// Normalize weights
		s.normalizeWeights()

		// Reset popularity counters
		s.postPopularity = make(map[string]int)
		mutex.Unlock()
	}
}

func (s *SimulationController) normalizeWeights() {
	total := 0.0
	for _, weight := range s.subredditWeights {
		total += weight
	}

	for subreddit := range s.subredditWeights {
		s.subredditWeights[subreddit] /= total
	}
}

func (s *SimulationController) simulateConnections() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		connectedCount := 0
		hour := time.Now().Hour()

		for _, client := range s.clients {
			baseProb := 0.8
			if hour >= 22 || hour <= 6 {
				baseProb = 0.3
			} else if hour >= 9 && hour <= 17 {
				baseProb = 0.9
			}

			connected := rand.Float32() < float32(baseProb)
			if connected {
				connectedCount++
			}
			s.system.Root.Send(client, &common.ConnectionStatus{Connected: connected})
		}

		connectionRate := float64(connectedCount) / float64(len(s.clients))
		s.metrics.UpdateSimulationMetrics(float64(len(s.clients)), connectionRate)
	}
}

func (s *SimulationController) reportMetrics() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		s.metrics.ReportCurrentStats()
	}
}

func (s *SimulationController) RunLoadTest(duration time.Duration, userIncrement int) {
	log.Printf("Starting load test: adding %d users every 5 seconds for %v", userIncrement, duration)
	initialUsers := len(s.clients)
	ticker := time.NewTicker(5 * time.Second)
	deadline := time.Now().Add(duration)
	defer ticker.Stop()

	for time.Now().Before(deadline) {
		select {
		case <-ticker.C:
			currentUsers := len(s.clients)
			log.Printf("Adding %d users. Current total: %d", userIncrement, currentUsers)
			err := s.addUsers(userIncrement)
			if err != nil {
				log.Printf("Failed to add users: %v", err)
				continue
			}
			s.metrics.UpdateActiveUsers(float64(len(s.clients)))
			// Record metrics without trying to get error count directly
			s.metrics.RecordLoadTestMetrics(
				float64(len(s.clients)),
				float64(currentUsers-initialUsers),
				0, // Remove direct counter access
			)
		}
	}
}
func (s *SimulationController) addUsers(count int) error {
	startIndex := len(s.clients)

	for i := 0; i < count; i++ {
		behavior := common.GenerateUserBehavior()

		clientActor := actor.NewClientActor(
			utils.GenerateID(),
			utils.GenerateUsername(startIndex+i),
			s.enginePID,
			behavior,
			s.metrics,
		)

		props := protoactor.PropsFromProducer(func() protoactor.Actor {
			return clientActor
		})

		pid, err := s.system.Root.SpawnNamed(props, fmt.Sprintf("client-%d", startIndex+i))
		if err != nil {
			return fmt.Errorf("failed to spawn client actor: %v", err)
		}

		s.clients = append(s.clients, pid)
	}

	return nil
}

func (s *SimulationController) SimulateNetworkConditions(packetLossRate float64) {
	log.Printf("Simulating network conditions with %f packet loss rate", packetLossRate)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		for _, client := range s.clients {
			if rand.Float64() < packetLossRate {
				// Simulate disconnection
				s.system.Root.Send(client, &common.ConnectionStatus{Connected: false})
			} else {
				// Simulate reconnection
				s.system.Root.Send(client, &common.ConnectionStatus{Connected: true})
			}
		}
	}
}
