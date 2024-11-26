package simulation

import (
	"fmt"
	protoactor "github.com/asynkron/protoactor-go/actor"
	"math/rand"
	"reddit-clone/internal/actor"
	"reddit-clone/internal/common"
	"reddit-clone/pkg/metrics"
	"reddit-clone/pkg/utils"
	"time"
)

type SimulationController struct {
	system    *protoactor.ActorSystem
	enginePID *protoactor.PID
	clients   []*protoactor.PID
	zipf      *rand.Zipf
	metrics   *metrics.RedditMetrics
}

func NewSimulationController(system *protoactor.ActorSystem, enginePID *protoactor.PID, metrics *metrics.RedditMetrics) *SimulationController {
	// Create a new random source
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Initialize Zipf with appropriate parameters
	zipf := rand.NewZipf(r, 1.1, 1.0, 1000)
	return &SimulationController{
		system:    system,
		enginePID: enginePID,
		clients:   make([]*protoactor.PID, 0),
		metrics:   metrics,
		zipf:      zipf,
	}
}

func (s *SimulationController) Start(numClients int) error {
	if s.enginePID == nil {
		return fmt.Errorf("engine PID is nil")
	}

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	s.zipf = rand.NewZipf(r, 1.1, 1.0, uint64(numClients))

	personaCounts := make(map[string]int)

	for i := 0; i < numClients; i++ {
		behavior := GenerateUserBehavior()
		personaCounts[behavior.Persona]++

		clientActor := actor.NewClientActor(
			utils.GenerateID(),
			utils.GenerateUsername(i),
			s.enginePID,
			behavior,
			s.metrics,
		)

		props := protoactor.PropsFromProducer(func() protoactor.Actor {
			return clientActor
		})

		pid, err := s.system.Root.SpawnNamed(props, fmt.Sprintf("client-%d", i))
		if err != nil {
			return fmt.Errorf("failed to spawn client actor: %v", err)
		}
		s.clients = append(s.clients, pid)
	}

	for persona, count := range personaCounts {
		s.metrics.UpdatePersonaCount(persona, count)
	}

	go s.simulateActions()
	go s.simulateConnections()
	go s.reportMetrics()

	return nil
}

func (s *SimulationController) simulateActions() {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		clientIndex := s.zipf.Uint64()
		if int(clientIndex) < len(s.clients) {
			s.system.Root.Send(s.clients[clientIndex], &common.SimulateAction{
				Timestamp: time.Now(),
			})
		}
	}
}

func (s *SimulationController) simulateConnections() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		connectedCount := 0
		for _, client := range s.clients {
			connected := rand.Float32() > 0.2 // 80% chance of being connected
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
