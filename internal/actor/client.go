package actor

import (
	"fmt"
	protoactor "github.com/asynkron/protoactor-go/actor"
	"math/rand"
	"reddit-clone/api/proto/generated"
	pb "reddit-clone/api/proto/generated"
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
	// Generate unique name using timestamp and user ID
	uniqueName := fmt.Sprintf("user-%s-%d", userID, time.Now().UnixNano())
	return &ClientActor{
		userID:     userID,
		username:   uniqueName,
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
	case *pb.PingMessage:
		context.Respond(&pb.PongMessage{})
	case *common.SimulateAction:
		if c.connected {
			start := time.Now()

			action := c.performAction(context)
			switch actionMsg := action.(interface{}).(type) {
			case *generated.PostMessage:
				future := context.RequestFuture(c.enginePID, actionMsg, 5*time.Second)
				if _, err := future.Result(); err == nil {
					c.metrics.UpdateActiveUsers(1)
					c.metrics.RecordAction(c.persona, "post")
				} else {
					c.metrics.RecordError()
				}

			case *generated.CommentMessage:
				future := context.RequestFuture(c.enginePID, actionMsg, 5*time.Second)
				if _, err := future.Result(); err == nil {
					c.metrics.UpdateActiveUsers(1)
					c.metrics.RecordAction(c.persona, "comment")
				} else {
					c.metrics.RecordError()
				}

			case *generated.VoteMessage:
				future := context.RequestFuture(c.enginePID, actionMsg, 5*time.Second)
				if _, err := future.Result(); err == nil {
					c.metrics.UpdateActiveUsers(1)
					c.metrics.RecordAction(c.persona, "vote")
				} else {
					c.metrics.RecordError()
				}
			}

			duration := time.Since(start).Seconds()
			c.metrics.RecordSimulatedAction(duration)
		}
	//case *common.SimulateAction:
	//	if c.connected {
	//		start := time.Now()
	//
	//		// Perform action and record metrics
	//		action := c.performAction(context)
	//		//if action != nil {
	//		// Send action to engine and wait for response
	//
	//		switch actionMsg := action.(interface{}).(type) {
	//		case *generated.PostMessage,
	//			*generated.CommentMessage,
	//			*generated.VoteMessage,
	//			*generated.JoinSubredditMessage:
	//
	//			future := context.RequestFuture(c.enginePID, action, 5*time.Second)
	//			if _, err := future.Result(); err == nil {
	//				// Update metrics based on successful action
	//				c.metrics.UpdateActiveUsers(1)
	//				duration := time.Since(start).Seconds()
	//				c.metrics.RecordSimulatedAction(duration)
	//			} else {
	//				c.metrics.RecordError()
	//			}
	//		}
	//	}

	case *common.ConnectionStatus:
		wasConnected := c.connected
		c.connected = msg.Connected

		// Update active users count
		if c.connected && !wasConnected {
			c.metrics.UpdateActiveUsers(1)
		} else if !c.connected && wasConnected {
			c.metrics.UpdateActiveUsers(-1)
		}
	}
}

//	// Send action to engine
//	response, err := context.RequestFuture(c.enginePID, action, 5*time.Second).Result()
//	if err != nil {
//		c.metrics.RecordError()
//		return
//	}
//
//	// Record timing and metrics
//	duration := time.Since(start).Seconds()
//	c.metrics.RecordSimulatedAction(duration)
//	c.metrics.RecordAction(c.persona, msg.ActionType)
//
//	// Update metrics based on response
//	if resp, ok := response.(*common.ActionResponse); ok {
//		c.metrics.UpdateActiveUsers(float64(resp.ActiveUsers))
//	}
//}

//	case *common.ConnectionStatus:
//		c.connected = msg.Connected
//		if c.connected {
//			c.metrics.UpdateActiveUsers(1)
//		} else {
//			c.metrics.UpdateActiveUsers(-1)
//		}
//	}
//}

func (c *ClientActor) performAction(context protoactor.Context) interface{} {
	//start := time.Now()

	if !c.isActiveHour(time.Now().Hour()) {
		return &pb.EmptyMessage{}
	}
	//var actionType common.ActionType
	rand := rand.Float64()
	switch {
	case rand < c.behavior.PostProbability:
		//actionType = common.PostAction
		//c.createPost(context)
		c.metrics.RecordAction(c.persona, "post")
		return c.createPost(context)
	case rand < c.behavior.PostProbability+c.behavior.CommentProbability:
		//actionType = common.CommentAction
		//c.createComment(context)
		c.metrics.RecordAction(c.persona, "comment")
		return c.createComment(context)
	case rand < c.behavior.PostProbability+c.behavior.CommentProbability+c.behavior.VoteProbability:
		//actionType = common.VoteAction
		//c.vote(context)
		c.metrics.RecordAction(c.persona, "vote")
		return c.vote(context)
	default:
		//actionType = common.JoinAction
		//c.joinSubreddit(context)

		c.metrics.RecordAction(c.persona, "join")
		return c.joinSubreddit(context)
	}

}

// c.metrics.RecordSimulatedAction(time.Since(start).Seconds())
// }
// made change here
func (c *ClientActor) createPost(context protoactor.Context) *generated.PostMessage {
	if len(c.subreddits) == 0 {
		return nil
	}

	subreddit := c.distribution.GetRandomSubreddit()
	if !c.distribution.ShouldCreatePost(subreddit) {
		return nil
	}

	content := ""
	isRepost := false
	if rand.Float64() < 0.2 { // 20% chance of repost
		content = c.getRandomExistingPost()
		isRepost = true
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
		IsRepost:    isRepost,
	}

	context.Request(c.enginePID, post)
	return post
}

func (c *ClientActor) createComment(context protoactor.Context) *generated.CommentMessage {
	if len(c.subreddits) == 0 {
		return nil
	}

	comment := &generated.CommentMessage{
		Id:        utils.GenerateID(),
		PostId:    utils.GenerateID(), // In real implementation, get actual post ID
		ParentId:  "",                 // Root level comment
		AuthorId:  c.userID,
		Content:   utils.GenerateRandomContent(),
		CreatedAt: time.Now().Unix(),
	}

	context.Request(c.enginePID, comment)
	return comment
}

func (c *ClientActor) joinSubreddit(context protoactor.Context) *generated.JoinSubredditMessage {
	if len(c.distribution.GetSubreddits()) == 0 {
		return nil
	}

	subredditID := c.distribution.GetRandomSubreddit()
	if subredditID == "" {
		return nil
	}

	join := &generated.JoinSubredditMessage{
		SubredditId: subredditID,
		UserId:      c.userID,
	}

	context.Request(c.enginePID, join)
	return join
}

func (c *ClientActor) vote(context protoactor.Context) *generated.VoteMessage {
	vote := &generated.VoteMessage{
		TargetId: utils.GenerateID(), // In real implementation, get actual post/comment ID
		UserId:   c.userID,
		IsUpvote: rand.Float32() > 0.3, // 70% chance of upvote
	}

	context.Request(c.enginePID, vote)
	return vote
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
