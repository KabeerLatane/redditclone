package actor

import (
	"github.com/asynkron/protoactor-go/actor"
	pb "reddit-clone/api/proto/generated" // Use pb as alias for generated code
	"reddit-clone/internal/models"
	"reddit-clone/internal/store"
	"reddit-clone/pkg/metrics"
	"time"
)

type EngineActor struct {
	store   store.Store
	metrics *metrics.RedditMetrics
}

func NewEngineActor(store store.Store, metrics *metrics.RedditMetrics) *EngineActor {
	return &EngineActor{
		store:   store,
		metrics: metrics,
	}
}

func (e *EngineActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *pb.UserMessage:
		e.handleUserMessage(context, msg)
	case *pb.SubredditMessage:
		e.handleSubredditMessage(context, msg)
	case *pb.PostMessage:
		e.handlePostMessage(context, msg)
	case *pb.VoteMessage:
		e.handleVoteMessage(context, msg)
	}
}

func (e *EngineActor) handleUserMessage(context actor.Context, msg *pb.UserMessage) {
	start := time.Now()
	user := &models.User{
		ID:       msg.UserId, // Changed from UserID
		Username: msg.Username,
		Password: msg.Password,
		Created:  time.Now().Unix(),
	}

	err := e.store.CreateUser(user)
	if err != nil {
		e.metrics.RecordError()
		context.Respond(&pb.ErrorResponse{Error: err.Error()})
		return
	}

	e.metrics.TotalUsers.Inc()
	e.metrics.RecordRequest(time.Since(start).Seconds())
	context.Respond(&pb.SuccessResponse{Message: "User registered successfully"})
}

func (e *EngineActor) handleSubredditMessage(context actor.Context, msg *pb.SubredditMessage) {
	start := time.Now()

	subreddit := &models.Subreddit{
		ID:          msg.Id,
		Name:        msg.Name,
		Description: msg.Description,
		CreatorID:   msg.CreatorId,
		Members:     make(map[string]bool),
		Created:     time.Now().Unix(),
	}

	err := e.store.CreateSubreddit(subreddit)
	if err != nil {
		e.metrics.RecordError()
		context.Respond(&pb.ErrorResponse{Error: err.Error()})
		return
	}

	e.metrics.UpdateSubredditMembers(subreddit.Name, 1) // Creator is first member
	e.metrics.RecordRequest(time.Since(start).Seconds())
	context.Respond(&pb.SuccessResponse{Message: "Subreddit created successfully"})
}

func (e *EngineActor) handlePostMessage(context actor.Context, msg *pb.PostMessage) {
	post := &models.Post{
		ID:          msg.Id,
		SubredditID: msg.SubredditId,
		AuthorID:    msg.AuthorId,
		Title:       msg.Title,
		Content:     msg.Content,
		Created:     time.Now().Unix(),
		Votes:       make(map[string]bool),
	}

	err := e.store.CreatePost(post)
	if err != nil {
		context.Respond(&pb.ErrorResponse{Error: err.Error()})
		return
	}

	context.Respond(&pb.SuccessResponse{Message: "Post created successfully"})
}

func (e *EngineActor) handleVoteMessage(context actor.Context, msg *pb.VoteMessage) {
	err := e.store.Vote(msg.TargetId, msg.UserId, msg.IsUpvote)
	if err != nil {
		context.Respond(&pb.ErrorResponse{Error: err.Error()})
		return
	}

	context.Respond(&pb.SuccessResponse{Message: "Vote recorded successfully"})
}
