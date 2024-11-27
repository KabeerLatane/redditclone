package main

import (
	"fmt"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	//pb "reddit-clone/api/proto/generated"
	internalActor "reddit-clone/internal/actor" // Alias the import
	"reddit-clone/internal/store/memory"
	"reddit-clone/pkg/metrics"
)

func main() {
	// Initialize metrics
	metricsCollector := metrics.NewRedditMetrics()

	// Start metrics endpoint
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(":2112", nil); err != nil {
			log.Fatalf("Failed to start metrics server: %v", err)
		}
	}()

	// Initialize actor system
	system := actor.NewActorSystem()

	// Configure remote
	remoteConfig := remote.Configure("127.0.0.1", 8090)

	// Create remote
	remoting := remote.NewRemote(system, remoteConfig)

	// Create new engine actor
	engineActor := internalActor.NewEngineActor(
		memory.NewMemoryStore(),
		metricsCollector,
	)

	// Create props
	props := actor.PropsFromProducer(func() actor.Actor {
		return engineActor
	})

	remoting.Register("engine", props)
	remoting.Start()

	// Spawn the engine actor
	pid, err := system.Root.SpawnNamed(props, "engine")
	if err != nil {
		log.Fatalf("Failed to spawn engine actor: %v", err)
	}

	fmt.Printf("Engine started. PID: %v\n", pid)
	// Add signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Create shutdown channel
	done := make(chan bool)

	// Handle shutdown gracefully
	go func() {
		sig := <-sigChan
		log.Printf("Received signal: %v", sig)

		// Cleanup
		remoting.Shutdown(true)
		system.Shutdown()

		done <- true
	}()
	// Keep the process running
	select {}
}
