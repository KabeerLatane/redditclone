// pkg/metrics/prometheus.go
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// PrometheusConfig holds the prometheus configuration
type PrometheusConfig struct {
	Port string
}

// DefaultConfig returns default prometheus configuration
func DefaultConfig() *PrometheusConfig {
	return &PrometheusConfig{
		Port: ":2112",
	}
}

// MetricsCollector holds all prometheus metrics
type MetricsCollector struct {
	// User metrics
	ActiveUsers prometheus.Gauge
	TotalUsers  prometheus.Counter

	// Content metrics
	PostsCreated    prometheus.Counter
	CommentsCreated prometheus.Counter
	VotesRecorded   prometheus.Counter

	// Performance metrics
	RequestDuration prometheus.Histogram
	ErrorRate       prometheus.Counter

	// Subreddit metrics
	SubredditMembers *prometheus.GaugeVec
	SubredditPosts   *prometheus.CounterVec

	// Simulation metrics
	SimulatedUsers    prometheus.Gauge
	ConnectionRate    prometheus.Gauge
	SimulatedActions  prometheus.Counter
	SimulationLatency prometheus.Histogram
}

// NewMetricsCollector initializes and registers all prometheus metrics
func NewMetricsCollector() *MetricsCollector {
	return &MetricsCollector{
		ActiveUsers: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "reddit_active_users",
			Help: "Number of currently active users",
		}),

		TotalUsers: promauto.NewCounter(prometheus.CounterOpts{
			Name: "reddit_total_users",
			Help: "Total number of registered users",
		}),

		PostsCreated: promauto.NewCounter(prometheus.CounterOpts{
			Name: "reddit_posts_total",
			Help: "Total number of posts created",
		}),

		CommentsCreated: promauto.NewCounter(prometheus.CounterOpts{
			Name: "reddit_comments_total",
			Help: "Total number of comments created",
		}),

		VotesRecorded: promauto.NewCounter(prometheus.CounterOpts{
			Name: "reddit_votes_total",
			Help: "Total number of votes recorded",
		}),

		RequestDuration: promauto.NewHistogram(prometheus.HistogramOpts{
			Name:    "reddit_request_duration_seconds",
			Help:    "Duration of requests in seconds",
			Buckets: prometheus.DefBuckets,
		}),

		ErrorRate: promauto.NewCounter(prometheus.CounterOpts{
			Name: "reddit_errors_total",
			Help: "Total number of errors encountered",
		}),

		SubredditMembers: promauto.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "reddit_subreddit_members",
				Help: "Number of members per subreddit",
			},
			[]string{"subreddit"},
		),

		SubredditPosts: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "reddit_subreddit_posts",
				Help: "Number of posts per subreddit",
			},
			[]string{"subreddit"},
		),

		SimulatedUsers: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "reddit_simulated_users",
			Help: "Number of simulated users",
		}),

		ConnectionRate: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "reddit_connection_rate",
			Help: "Rate of connected users",
		}),

		SimulatedActions: promauto.NewCounter(prometheus.CounterOpts{
			Name: "reddit_simulated_actions_total",
			Help: "Total number of simulated actions",
		}),

		SimulationLatency: promauto.NewHistogram(prometheus.HistogramOpts{
			Name:    "reddit_simulation_latency_seconds",
			Help:    "Latency of simulated actions in seconds",
			Buckets: prometheus.DefBuckets,
		}),
	}
}

// RecordRequest records the duration of a request
func (m *MetricsCollector) RecordRequest(duration float64) {
	m.RequestDuration.Observe(duration)
}

// RecordError increments the error counter
func (m *MetricsCollector) RecordError() {
	m.ErrorRate.Inc()
}

// UpdateSubredditMembers updates the member count for a subreddit
func (m *MetricsCollector) UpdateSubredditMembers(subreddit string, count float64) {
	m.SubredditMembers.WithLabelValues(subreddit).Set(count)
}

// IncrementSubredditPosts increments the post count for a subreddit
func (m *MetricsCollector) IncrementSubredditPosts(subreddit string) {
	m.SubredditPosts.WithLabelValues(subreddit).Inc()
}

// UpdateSimulationMetrics updates simulation-related metrics
func (m *MetricsCollector) UpdateSimulationMetrics(users float64, connectionRate float64) {
	m.SimulatedUsers.Set(users)
	m.ConnectionRate.Set(connectionRate)
}

// RecordSimulatedAction records a simulated action and its latency
func (m *MetricsCollector) RecordSimulatedAction(latency float64) {
	m.SimulatedActions.Inc()
	m.SimulationLatency.Observe(latency)
}
