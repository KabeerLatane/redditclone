// pkg/metrics/collector.go
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	dto "github.com/prometheus/client_model/go"
	"net/http"
	"sync"
)

// RedditMetrics holds all metrics for the Reddit clone
type RedditMetrics struct {
	PostsCreated        prometheus.Counter
	CommentsCreated     prometheus.Counter
	VotesRecorded       prometheus.Counter
	ActiveUsers         prometheus.Gauge
	ResponseTime        prometheus.Histogram
	SubredditSize       *prometheus.GaugeVec
	TotalUsers          prometheus.Counter
	ErrorCount          prometheus.Counter
	PersonaStats        map[string]*PersonaStats
	SimulatedUsers      prometheus.Gauge
	AverageResponseTime prometheus.Gauge
	ErrorRate           prometheus.Gauge
}

type PersonaStats struct {
	ActiveUsers  int
	PostCount    int
	VoteCount    int
	CommentCount int
}

var (
	metricsOnce     sync.Once
	metricsInstance *RedditMetrics
)

// NewRedditMetrics creates a new Reddit metrics collector
func NewRedditMetrics() *RedditMetrics {
	metricsOnce.Do(func() {
		metricsInstance = &RedditMetrics{
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
			SimulatedUsers: promauto.NewGauge(prometheus.GaugeOpts{
				Name: "reddit_simulated_users_total",
				Help: "Total number of simulated users",
			}),
			AverageResponseTime: promauto.NewGauge(prometheus.GaugeOpts{
				Name: "reddit_average_response_time_seconds",
				Help: "Average response time in seconds",
			}),
			ErrorRate: promauto.NewGauge(prometheus.GaugeOpts{
				Name: "reddit_error_rate",
				Help: "Rate of errors per second",
			}),
		}

	})
	return metricsInstance
}

//delete this later
//func (m *RedditMetrics) Get() map[string]float64 {
//	metrics := make(map[string]float64)
//
//	// Get values using prometheus.Gauge methods
//	metrics["active_users"] = getGaugeValue(m.ActiveUsers)
//	metrics["simulated_users"] = getGaugeValue(m.SimulatedUsers)
//
//	// Get values using prometheus.Counter methods
//	metrics["total_users"] = getCounterValue(m.TotalUsers)
//	metrics["posts"] = getCounterValue(m.PostsCreated)
//	metrics["comments"] = getCounterValue(m.CommentsCreated)
//	metrics["votes"] = getCounterValue(m.VotesRecorded)
//
//	return metrics
//}
//
//// Helper functions to get values from different metric types
//func getGaugeValue(g prometheus.Gauge) float64 {
//	var m dto.Metric
//	g.Write(&m)
//	return m.GetGauge().GetValue()
//}
//
//func getCounterValue(c prometheus.Counter) float64 {
//	var m dto.Metric
//	c.Write(&m)
//	return m.GetCounter().GetValue()
//}

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
	m.AverageResponseTime.Set(duration)
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
func (m *RedditMetrics) RecordLoadTestMetrics(users, responseTime, errorRate float64) {
	m.SimulatedUsers.Set(users)
	m.AverageResponseTime.Set(responseTime)
	m.ErrorRate.Set(errorRate)
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
func (m *RedditMetrics) UpdateActiveUsers(count float64) {
	m.ActiveUsers.Set(count)
	m.SimulatedUsers.Set(count)
	m.TotalUsers.Add(count)
}

func (m *RedditMetrics) GetEssentialMetrics() map[string]float64 {
	metrics := make(map[string]float64)

	// Get current values directly using prometheus.Gauge methods
	var metric dto.Metric

	// Get Gauge values
	m.ActiveUsers.Write(&metric)
	metrics["active_users"] = metric.GetGauge().GetValue()

	m.SimulatedUsers.Write(&metric)
	metrics["simulated_users"] = metric.GetGauge().GetValue()

	// Get Counter values
	m.PostsCreated.Write(&metric)
	metrics["posts"] = metric.GetCounter().GetValue()

	m.CommentsCreated.Write(&metric)
	metrics["comments"] = metric.GetCounter().GetValue()

	m.VotesRecorded.Write(&metric)
	metrics["votes"] = metric.GetCounter().GetValue()

	m.TotalUsers.Write(&metric)
	metrics["total_users"] = metric.GetCounter().GetValue()

	// Get other gauge values
	m.AverageResponseTime.Write(&metric)
	metrics["request_duration"] = metric.GetGauge().GetValue()

	m.ErrorRate.Write(&metric)
	metrics["error_rate"] = metric.GetGauge().GetValue()

	return metrics
}
