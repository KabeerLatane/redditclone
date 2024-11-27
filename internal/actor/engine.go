package actor

import (
	"github.com/asynkron/protoactor-go/actor"
	pb "reddit-clone/api/proto/generated"
	"reddit-clone/internal/models"
	"reddit-clone/internal/store"
	"reddit-clone/pkg/metrics"
	"sort"
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
	case *pb.PingMessage:
		context.Respond(&pb.PongMessage{})
	case *pb.UserMessage:
		e.handleUserMessage(context, msg)
	case *pb.SubredditMessage:
		e.handleSubredditMessage(context, msg)
	case *pb.JoinSubredditMessage:
		e.handleJoinSubredditMessage(context, msg)
	case *pb.PostMessage:
		e.handlePostMessage(context, msg)
	case *pb.CommentMessage:
		e.handleCommentMessage(context, msg)
	case *pb.VoteMessage:
		e.handleVoteMessage(context, msg)
	case *pb.DirectMessageMessage:
		e.handleDirectMessage(context, msg)
	case *pb.GetFeedMessage:
		e.handleGetFeed(context, msg)
	case *pb.GetCommentsMessage:
		e.handleGetComments(context, msg)
	}
}

func (e *EngineActor) handleJoinSubredditMessage(context actor.Context, msg *pb.JoinSubredditMessage) {
	start := time.Now()

	err := e.store.JoinSubreddit(msg.SubredditId, msg.UserId)
	if err != nil {
		e.metrics.RecordError()
		context.Respond(&pb.ErrorResponse{Error: err.Error()})
		return
	}

	e.metrics.UpdateSubredditMembers(msg.SubredditId, 1)
	e.metrics.RecordRequest(time.Since(start).Seconds())
	context.Respond(&pb.SuccessResponse{Message: "Joined subreddit successfully"})
}

func (e *EngineActor) handleUserMessage(context actor.Context, msg *pb.UserMessage) {
	start := time.Now()
	user := &models.User{
		ID:       msg.UserId,
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

	e.metrics.UpdateSubredditMembers(subreddit.Name, 1)
	e.metrics.RecordRequest(time.Since(start).Seconds())
	context.Respond(&pb.SuccessResponse{Message: "Subreddit created successfully"})
}

func (e *EngineActor) handlePostMessage(context actor.Context, msg *pb.PostMessage) {
	start := time.Now()

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
		e.metrics.RecordError()
		context.Respond(&pb.ErrorResponse{Error: err.Error()})
		return
	}

	e.metrics.PostsCreated.Inc()
	e.metrics.RecordRequest(time.Since(start).Seconds())
	context.Respond(&pb.SuccessResponse{Message: "Post created successfully"})
}

func (e *EngineActor) handleCommentMessage(context actor.Context, msg *pb.CommentMessage) {
	start := time.Now()

	comment := &models.Comment{
		ID:       msg.Id,
		PostID:   msg.PostId,
		ParentID: msg.ParentId,
		AuthorID: msg.AuthorId,
		Content:  msg.Content,
		Created:  time.Now().Unix(),
	}

	err := e.store.AddComment(comment)
	if err != nil {
		e.metrics.RecordError()
		context.Respond(&pb.ErrorResponse{Error: err.Error()})
		return
	}

	e.metrics.CommentsCreated.Inc()
	e.metrics.RecordRequest(time.Since(start).Seconds())
	context.Respond(&pb.SuccessResponse{Message: "Comment created successfully"})
}

func (e *EngineActor) handleVoteMessage(context actor.Context, msg *pb.VoteMessage) {
	start := time.Now()

	err := e.store.Vote(msg.TargetId, msg.UserId, msg.IsUpvote)
	if err != nil {
		e.metrics.RecordError()
		context.Respond(&pb.ErrorResponse{Error: err.Error()})
		return
	}

	e.metrics.VotesRecorded.Inc()
	e.metrics.RecordRequest(time.Since(start).Seconds())
	context.Respond(&pb.SuccessResponse{Message: "Vote recorded successfully"})
}

func (e *EngineActor) handleDirectMessage(context actor.Context, msg *pb.DirectMessageMessage) {
	start := time.Now()

	message := &models.DirectMessage{
		ID:        msg.GetId(),      // Use GetId() method
		FromID:    msg.GetFromId(),  // Use GetFromId() method
		ToID:      msg.GetToId(),    // Use GetToId() method
		Content:   msg.GetContent(), // Use GetContent() method
		Timestamp: time.Now().Unix(),
	}

	err := e.store.SendMessage(message)
	if err != nil {
		e.metrics.RecordError()
		context.Respond(&pb.ErrorResponse{Error: err.Error()})
		return
	}

	e.metrics.RecordRequest(time.Since(start).Seconds())
	context.Respond(&pb.SuccessResponse{Message: "Message sent successfully"})
}

func (e *EngineActor) handleGetFeed(context actor.Context, msg *pb.GetFeedMessage) {
	start := time.Now()

	// Get posts from subscribed subreddits
	var feed []*models.Post
	for _, subredditID := range msg.SubredditIds {
		posts, err := e.store.GetSubredditPosts(subredditID)
		if err != nil {
			e.metrics.RecordError()
			context.Respond(&pb.ErrorResponse{Error: err.Error()})
			return
		}
		feed = append(feed, posts...)
	}

	// Sort by creation time and karma
	sort.Slice(feed, func(i, j int) bool {
		scoreI := float64(feed[i].Karma) / time.Since(time.Unix(feed[i].Created, 0)).Hours()
		scoreJ := float64(feed[j].Karma) / time.Since(time.Unix(feed[j].Created, 0)).Hours()
		return scoreI > scoreJ
	})

	// Convert to proto message
	response := &pb.FeedResponse{
		Posts: make([]*pb.PostMessage, 0, len(feed)),
	}
	for _, post := range feed {
		response.Posts = append(response.Posts, &pb.PostMessage{
			Id:          post.ID,
			SubredditId: post.SubredditID,
			AuthorId:    post.AuthorID,
			Title:       post.Title,
			Content:     post.Content,
			CreatedAt:   post.Created,
			IsRepost:    false,
		})
	}

	e.metrics.RecordRequest(time.Since(start).Seconds())
	context.Respond(response)
}

func (e *EngineActor) handleGetComments(context actor.Context, msg *pb.GetCommentsMessage) {
	start := time.Now()

	comments, err := e.store.GetComments(msg.PostId)
	if err != nil {
		e.metrics.RecordError()
		context.Respond(&pb.ErrorResponse{Error: err.Error()})
		return
	}

	// Build comment tree
	commentMap := make(map[string]*models.Comment)
	rootComments := make([]*models.Comment, 0)

	for _, comment := range comments {
		commentMap[comment.ID] = comment
		if comment.ParentID == "" {
			rootComments = append(rootComments, comment)
		} else {
			if parent, exists := commentMap[comment.ParentID]; exists {
				parent.Children = append(parent.Children, comment.ID)
			}
		}
	}

	// Convert to proto message
	response := &pb.CommentsResponse{
		Comments: make([]*pb.CommentMessage, 0, len(comments)),
	}

	// Helper function to convert comment tree to proto
	var convertComment func(*models.Comment) *pb.CommentMessage
	convertComment = func(comment *models.Comment) *pb.CommentMessage {
		protoComment := &pb.CommentMessage{
			Id:        comment.ID,
			PostId:    comment.PostID,
			ParentId:  comment.ParentID,
			AuthorId:  comment.AuthorID,
			Content:   comment.Content,
			CreatedAt: comment.Created,
		}
		return protoComment
	}

	// Convert all comments starting from root
	for _, rootComment := range rootComments {
		response.Comments = append(response.Comments, convertComment(rootComment))
	}

	e.metrics.RecordRequest(time.Since(start).Seconds())
	context.Respond(response)
}
