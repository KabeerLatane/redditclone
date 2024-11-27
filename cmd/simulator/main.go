// cmd/simulator/main.go
package main

import (
	"fmt"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	pb "reddit-clone/api/proto/generated"
	"reddit-clone/internal/simulation"
	"reddit-clone/pkg/metrics"
	"strings"
	"time"
)

func main() {
	// Initialize metrics collector
	metricsCollector := metrics.NewRedditMetrics()

	// Start metrics endpoint
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(":2113", nil); err != nil {
			log.Fatalf("Failed to start metrics server: %v", err)
		}
	}()

	// Initialize actor system
	system := actor.NewActorSystem()

	// Configure remote
	remoteConfig := remote.Configure("127.0.0.1", 8091)
	remoting := remote.NewRemote(system, remoteConfig)
	remoting.Start()

	// Wait for engine to be available
	time.Sleep(10 * time.Second)

	//// Connect to engine
	//enginePID := actor.NewPID("localhost:8090", "engine")
	// Create engine PID
	enginePID := actor.NewPID("127.0.0.1:8090", "engine")
	log.Printf("Initializing connection to engine at %s...", enginePID.Address)
	// Verify engine connection

	timeout := 30 * time.Second
	future := system.Root.RequestFuture(enginePID, &pb.PingMessage{}, timeout) //here
	// Try to connect multiple times
	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		err := verifyEngineConnection(system, enginePID, 10*time.Second)
		if err == nil {
			log.Println("Successfully connected to engine")
			break
		}

		if i < maxRetries-1 {
			log.Printf("Connection attempt %d failed: %v. Retrying...", i+1, err)
			time.Sleep(2 * time.Second)
		} else {
			log.Fatalf("Failed to connect after %d attempts: %v", maxRetries, err)
		}
	}

	if _, err := future.Result(); err != nil {
		log.Fatalf("Failed to connect to engine: %v", err)
	}

	log.Println("Successfully connected to engine")
	// Create simulation controller
	controller := simulation.NewSimulationController(system, enginePID, metricsCollector)

	// Start simulation with 1000 clients
	if err := controller.Start(10); err != nil {
		log.Fatalf("Failed to start simulation: %v", err)
	}
	// Run test scenarios
	go runTestScenarios(controller)
	// Run indefinitely
	select {}
}

//	func verifyEngineConnection(system *actor.ActorSystem, enginePID *actor.PID, timeout time.Duration) error {
//		context := system.Root
//		// Use a longer timeout for initial connection
//		future := context.RequestFuture(enginePID, &pb.PingMessage{}, 10*time.Second)
//		_, err := future.Result()
//		return err
//	}
func verifyEngineConnection(system *actor.ActorSystem, enginePID *actor.PID, timeout time.Duration) error {
	log.Printf("Attempting to connect to engine at %s...", enginePID.Address)

	context := system.Root
	future := context.RequestFuture(enginePID, &pb.PingMessage{}, timeout)

	result, err := future.Result()
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "future: timeout"):
			log.Fatalf("Connection timed out after %v: engine might be down or unreachable", timeout)
		case strings.Contains(err.Error(), "msg must be proto.Message"):
			log.Fatalf("Message serialization failed: %v (check protobuf message definitions)", err)
		default:
			log.Fatalf("Connection failed: %v", err)
		}
	}
	if _, ok := result.(*pb.PongMessage); !ok {
		return fmt.Errorf("unexpected response type: %T (expected PongMessage)", result)
	}

	return nil
}

func runTestScenarios(controller *simulation.SimulationController) {
	// Wait for initial setup to stabilize
	time.Sleep(30 * time.Second)

	log.Println("Starting test scenarios...")

	// Scenario 1: Normal Load
	log.Println("Running normal load test...")
	controller.RunLoadTest(15*time.Minute, 10) // Add 10 users every 30 seconds for 15 minutes
	time.Sleep(1 * time.Minute)

	// Scenario 2: High Load
	log.Println("Running high load test...")
	controller.RunLoadTest(10*time.Minute, 50) // Add 50 users every 30 seconds for 10 minutes
	time.Sleep(1 * time.Minute)

	// Scenario 3: Network Issues
	log.Println("Testing network conditions...")
	go controller.SimulateNetworkConditions(0.2) // 20% packet loss
	time.Sleep(5 * time.Minute)

	// Scenario 4: Stress Test
	log.Println("Running stress test...")
	controller.RunLoadTest(5*time.Minute, 100) // Add 100 users every 30 seconds for 5 minutes

	log.Println("Test scenarios completed")
}
