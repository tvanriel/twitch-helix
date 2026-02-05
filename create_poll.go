package twitchhelix

import (
	"context"
	"time"
)

// RequestCreatePoll represents the request body used to create a poll on a broadcaster's channel.
type RequestCreatePoll struct {
	// BroadcasterID is the ID of the broadcaster creating the poll.
	BroadcasterID string `json:"broadcaster_id"`

	// Title is the question or title of the poll.
	Title string `json:"title"`

	// Choices is the list of options viewers can vote on.
	//
	// Must include at least 2 and at most 5 choices.
	Choices []Choice `json:"choices"`

	// ChannelPointsVotingEnabled enables voting with channel points.
	ChannelPointsVotingEnabled bool `json:"channel_points_voting_enabled"`

	// ChannelPointsPerVote is the cost in channel points per vote.
	//
	// Required if ChannelPointsVotingEnabled is true.
	ChannelPointsPerVote int `json:"channel_points_per_vote"`

	// DurationInSeconds specifies how long the poll will run in seconds.
	//
	// Minimum is 15 seconds, maximum is 1800 seconds (30 minutes).
	DurationInSeconds int `json:"duration"`
}

// ResponseCreatePoll represents the response returned after creating a poll.
type ResponseCreatePoll struct {
	ID                         string    `json:"id"`
	BroadcasterID              string    `json:"broadcaster_id"`
	BroadcasterName            string    `json:"broadcaster_name"`
	BroadcasterLogin           string    `json:"broadcaster_login"`
	Title                      string    `json:"title"`
	Choices                    []Choice  `json:"choices"`
	BitsVotingEnabled          bool      `json:"bits_voting_enabled"`
	BitsPerVote                int       `json:"bits_per_vote"`
	ChannelPointsVotingEnabled bool      `json:"channel_points_voting_enabled"`
	ChannelPointsPerVote       int       `json:"channel_points_per_vote"`
	Status                     string    `json:"status"` // "ACTIVE", "COMPLETED", or "TERMINATED"
	Duration                   int       `json:"duration"`
	StartedAt                  time.Time `json:"started_at"` // UTC timestamp when the poll started
}

// Choice represents an individual option in a poll.
type Choice struct {
	// ID is the unique identifier for this choice (set by Twitch on creation).
	ID string `json:"id,omitempty"`

	// Title is the text displayed for this choice.
	Title string `json:"title"`

	// Votes is the total number of votes this choice has received.
	Votes int `json:"votes,omitempty"`

	// ChannelPointsVotes is the total votes received from channel points.
	ChannelPointsVotes int `json:"channel_points_votes,omitempty"`

	// BitsVotes is the total votes received from bits.
	BitsVotes int `json:"bits_votes,omitempty"`
}

// CreatePoll creates a new poll on a broadcaster's channel.
//
// The broadcaster must be authenticated. Polls must have at least 2 choices and no more than 5.
// Optionally, viewers can vote using channel points. The poll will run for the specified duration.
func (c *Client) CreatePoll(ctx context.Context, req RequestCreatePoll) (*ResponseCreatePoll, error) {
	var resp ResponseCreatePoll

	err := c.doRequest(ctx, "POST", "polls", req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
