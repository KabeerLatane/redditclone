// pkg/metrics/collector.go
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// RedditMetrics holds all metrics for the Reddit clone
type RedditMetrics struct {
	PostsCreated    prometheus.Counter
	CommentsCreated prometheus.Counter
	VotesRecorded   prometheus.Counter
	ActiveUsers     prometheus.Gauge
	ResponseTime    prometheus.Histogram
	SubredditSize   *prometheus.GaugeVec
	TotalUsers      prometheus.Counter // Added this
	ErrorCount      prometheus.Counter // Added this
	PersonaStats    map[string]*PersonaStats
}

type PersonaStats struct {
	ActiveUsers  int
	PostCount    int
	VoteCount    int
	CommentCount int
}

// NewRedditMetrics creates a new Reddit metrics collector
func NewRedditMetrics() *RedditMetrics {
	return &RedditMetrics{
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
		ActiveUsers: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "reddit_active_users",
			Help: "Number of currently active users",
		}),
		ResponseTime: promauto.NewHistogram(prometheus.HistogramOpts{
			Name:    "reddit_request_duration_seconds",
			Help:    "Request duration in seconds",
			Buckets: prometheus.DefBuckets,
		}),
		SubredditSize: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Name: "reddit_subreddit_members",
			Help: "Number of members per subreddit",
		}, []string{"subreddit"}),
		TotalUsers: promauto.NewCounter(prometheus.CounterOpts{
			Name: "reddit_total_users",
			Help: "Total number of registered users",
		}),
		ErrorCount: promauto.NewCounter(prometheus.CounterOpts{
			Name: "reddit_errors_total",
			Help: "Total number of errors",
		}),
		PersonaStats: make(map[string]*PersonaStats),
	}
}

// RecordAction records metrics for different types of actions
func (m *RedditMetrics) RecordAction(persona string, actionType string) {
	if _, exists := m.PersonaStats[persona]; !exists {
		m.PersonaStats[persona] = &PersonaStats{}
	}

	switch actionType {
	case "post":
		m.PostsCreated.Inc()
		m.PersonaStats[persona].PostCount++
	case "comment":
		m.CommentsCreated.Inc()
		m.PersonaStats[persona].CommentCount++
	case "vote":
		m.VotesRecorded.Inc()
		m.PersonaStats[persona].VoteCount++
	}
}

// UpdatePersonaCount updates the active users count for a persona
func (m *RedditMetrics) UpdatePersonaCount(persona string, count int) {
	if _, exists := m.PersonaStats[persona]; !exists {
		m.PersonaStats[persona] = &PersonaStats{}
	}
	m.PersonaStats[persona].ActiveUsers = count
}

// UpdateSimulationMetrics updates general simulation metrics
func (m *RedditMetrics) UpdateSimulationMetrics(users, connectionRate float64) {
	m.ActiveUsers.Set(users)
}

// RecordSimulatedAction records the duration of a simulated action
func (m *RedditMetrics) RecordSimulatedAction(duration float64) {
	m.ResponseTime.Observe(duration)
}

// UpdateSubredditMembers updates the member count for a subreddit
func (m *RedditMetrics) UpdateSubredditMembers(subredditID string, count float64) {
	m.SubredditSize.WithLabelValues(subredditID).Set(count)
}

// RecordError increments the error counter
func (m *RedditMetrics) RecordError() {
	m.ErrorCount.Inc()
}

// RecordRequest records the duration of a request
func (m *RedditMetrics) RecordRequest(duration float64) {
	m.ResponseTime.Observe(duration)
}

// ReportCurrentStats generates a report of current statistics
func (m *RedditMetrics) ReportCurrentStats() {
	// Implementation for reporting current stats
	// This can be customized based on your needs
}
