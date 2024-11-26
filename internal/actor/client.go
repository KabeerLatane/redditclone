package actor

import (
	protoactor "github.com/asynkron/protoactor-go/actor"
	"math/rand"
	"reddit-clone/api/proto/generated"
	"reddit-clone/internal/common"
	"reddit-clone/pkg/metrics"
	"reddit-clone/pkg/utils"
	"time"
)

type ClientActor struct {
	userID       string
	username     string
	enginePID    *protoactor.PID
	connected    bool
	subreddits   []string
	metrics      *metrics.RedditMetrics
	behavior     *common.ClientBehavior
	distribution *common.SimulationDistribution
	persona      string
}

func NewClientActor(userID string, username string, enginePID *protoactor.PID, behavior *common.ClientBehavior, metrics *metrics.RedditMetrics) *ClientActor {
	return &ClientActor{
		userID:     userID,
		username:   username,
		enginePID:  enginePID,
		connected:  true,
		subreddits: make([]string, 0),
		metrics:    metrics,
		behavior:   behavior,
		persona:    behavior.Persona,
	}
}

func (c *ClientActor) Receive(context protoactor.Context) {
	switch msg := context.Message().(type) {
	case *common.SimulateAction:
		if c.connected {
			c.performRandomAction(context)
		}
	case *common.ConnectionStatus:
		c.connected = msg.Connected
	}
}

func (c *ClientActor) performRandomAction(context protoactor.Context) {
	hour := time.Now().Hour()
	if !c.isActiveHour(hour) {
		return
	}

	rand := rand.Float64()
	switch {
	case rand < c.behavior.PostProbability:
		c.createPost(context)
	case rand < c.behavior.PostProbability+c.behavior.CommentProbability:
		c.createComment(context)
	case rand < c.behavior.PostProbability+c.behavior.CommentProbability+c.behavior.VoteProbability:
		c.vote(context)
	default:
		c.joinSubreddit(context)
	}
}

func (c *ClientActor) createPost(context protoactor.Context) {
	if len(c.subreddits) == 0 {
		return
	}

	subreddit := c.distribution.GetRandomSubreddit()
	if !c.distribution.ShouldCreatePost(subreddit) {
		return
	}

	content := ""
	if rand.Float64() < 0.2 { // 20% chance of repost
		content = c.getRandomExistingPost()
	} else {
		content = utils.GenerateRandomContent()
	}

	post := &generated.PostMessage{
		Id:          utils.GenerateID(),
		SubredditId: subreddit,
		AuthorId:    c.userID,
		Title:       utils.GenerateRandomTitle(),
		Content:     content,
		CreatedAt:   time.Now().Unix(),
		IsRepost:    content != "",
	}

	start := time.Now()
	c.metrics.RecordSimulatedAction(time.Since(start).Seconds())
	context.Request(c.enginePID, post)
}

func (c *ClientActor) createComment(context protoactor.Context) {
	if len(c.subreddits) == 0 {
		return
	}

	comment := &generated.CommentMessage{
		Id:        utils.GenerateID(),
		PostId:    utils.GenerateID(), // In real implementation, get actual post ID
		ParentId:  "",                 // Root level comment
		AuthorId:  c.userID,
		Content:   utils.GenerateRandomContent(),
		CreatedAt: time.Now().Unix(),
	}

	start := time.Now()
	c.metrics.RecordSimulatedAction(time.Since(start).Seconds())
	context.Request(c.enginePID, comment)
}

func (c *ClientActor) joinSubreddit(context protoactor.Context) {
	if len(c.distribution.GetSubreddits()) == 0 {
		return
	}

	subredditID := c.distribution.GetRandomSubreddit()
	if subredditID == "" {
		return
	}

	join := &generated.JoinSubredditMessage{
		SubredditId: subredditID,
		UserId:      c.userID,
	}

	start := time.Now()
	c.metrics.RecordSimulatedAction(time.Since(start).Seconds())
	context.Request(c.enginePID, join)
}

func (c *ClientActor) vote(context protoactor.Context) {
	vote := &generated.VoteMessage{
		TargetId: utils.GenerateID(), // In real implementation, get actual post/comment ID
		UserId:   c.userID,
		IsUpvote: rand.Float32() > 0.3, // 70% chance of upvote
	}

	start := time.Now()
	c.metrics.RecordSimulatedAction(time.Since(start).Seconds())
	context.Request(c.enginePID, vote)
}

func (c *ClientActor) isActiveHour(hour int) bool {
	for _, activeHour := range c.behavior.ActiveHours {
		if hour == activeHour {
			return true
		}
	}
	return false
}

func (c *ClientActor) getRandomExistingPost() string {
	// TODO: Implement getting a random existing post
	return utils.GenerateRandomContent()
}

func (c *ClientActor) shouldPost() bool {
	return rand.Float64() < c.behavior.PostProbability
}

func (c *ClientActor) shouldComment() bool {
	return rand.Float64() < c.behavior.CommentProbability
}

func (c *ClientActor) shouldVote() bool {
	return rand.Float64() < c.behavior.VoteProbability
}

func (c *ClientActor) performAction(context protoactor.Context) {
	start := time.Now()

	if !c.isActiveHour(time.Now().Hour()) {
		return
	}

	switch {
	case c.shouldPost():
		c.createPost(context)
		c.metrics.RecordAction(c.persona, "post")
	case c.shouldComment():
		c.createComment(context)
		c.metrics.RecordAction(c.persona, "comment")
	case c.shouldVote():
		c.vote(context)
		c.metrics.RecordAction(c.persona, "vote")
	}

	c.metrics.RecordSimulatedAction(time.Since(start).Seconds())
}
