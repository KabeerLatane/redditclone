// cmd/engine/main.go
package main

import (
	"fmt"
	protoactor "github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"reddit-clone/internal/actor"
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

	// Initialize actor system with remote capability
	system := protoactor.NewActorSystem()
	remoteConfig := remote.Configure("localhost", 8090)
	remoting := remote.NewRemote(system, remoteConfig)
	remoting.Start()

	// Initialize store
	store := memory.NewMemoryStore()

	// Create engine actor with metrics
	engine := actor.NewEngineActor(store, metricsCollector)

	// Spawn the engine actor with a known ID
	props := protoactor.PropsFromProducer(func() protoactor.Actor {
		return engine
	})

	pid, err := system.Root.SpawnNamed(props, "engine")
	if err != nil {
		log.Fatalf("Failed to spawn engine actor: %v", err)
	}

	fmt.Printf("Engine started. PID: %v\n", pid)

	// Keep the process running
	select {}
}
