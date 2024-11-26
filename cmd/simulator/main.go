// cmd/simulator/main.go
package main

import (
	"fmt"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"reddit-clone/internal/simulation"
	"reddit-clone/pkg/metrics"
	"time"
)

func main() {
	// Initialize metrics collector
	metricsCollector := metrics.NewRedditMetrics()

	// Start metrics endpoint on a different port than the engine
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(":2113", nil); err != nil {
			log.Fatalf("Failed to start metrics server: %v", err)
		}
	}()

	// Initialize actor system with remote capability
	system := actor.NewActorSystem()
	remoteConfig := remote.Configure("localhost", 8091) // Different port than engine
	remoting := remote.NewRemote(system, remoteConfig)
	remoting.Start()

	// Wait for engine to be available
	time.Sleep(2 * time.Second)

	// Connect to engine using the correct address format
	enginePID := actor.NewPID("localhost:8090", "engine")

	// Verify engine connection
	timeout := 5 * time.Second
	err := verifyEngineConnection(system, enginePID, timeout)
	if err != nil {
		log.Fatalf("Failed to connect to engine: %v", err)
	}

	fmt.Printf("Connected to engine at %v\n", enginePID)

	// Create simulation controller
	controller := simulation.NewSimulationController(system, enginePID, metricsCollector)

	// Start simulation with 1000 clients
	if err := controller.Start(1000); err != nil {
		log.Fatalf("Failed to start simulation: %v", err)
	}

	// Run indefinitely
	select {}
}

func verifyEngineConnection(system *actor.ActorSystem, enginePID *actor.PID, timeout time.Duration) error {
	future := system.Root.RequestFuture(enginePID, &actor.PingRequest{}, timeout)
	_, err := future.Result()
	return err
}
