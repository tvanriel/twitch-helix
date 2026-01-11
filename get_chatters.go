package twitchhelix

import (
	"context"
	"fmt"

	"github.com/google/go-querystring/query"
)

// ChattersRequest represents the data required to request chatters connected to twitch chat.
type ChattersRequest struct {
	// BroadcasterID represents the id of the broadcaster whos chat you are checking.
	BroadcasterID string `url:"broadcaster_id"`

	// ModeratorID is the id of the moderator that is querying
	ModeratorID string `url:"moderator_id"`

	// First is the number of users you want on each page
	//
	// Default: 100
	// Maximum: 1000
	First *int `url:"first,omitempty"`

	// After represents the cursor to get the next page
	After *string `url:"after,omitempty"`
}

// ChattersResponse represents the information about the chatters and pagination information.
type ChattersResponse struct {
	// Chatters represents a slice of every chatter on that page.
	Chatters []Chatter `json:"data"`

	// Pagination contains the cursor for the next page
	//
	// May be empty if there are no pages left
	Pagination Pagination `json:"pagination"`

	// Total represents the total number of users connected to chat
	//
	// *NOT CCV
	Total int `json:"total"`
}

// Chatter represents basic user information
type Chatter struct {
	// UserID represents the id of the user
	UserID string `json:"user_id"`

	// UserLogin represents the login name of the user
	UserLogin string `json:"user_login"`

	// UserName represents the display name of the user
	UserName string `json:"user_name"`
}

// GetChatters gets the users connected to twitch's IRC chat client.
func (c *Client) GetChatters(ctx context.Context, req ChattersRequest) (*ChattersResponse, error) {
	var resp ChattersResponse
	values, err := query.Values(req)
	if err != nil {
		return nil, err
	}
	// Example Request URL:
	// https://api.twitch.tv/helix/chat.chatters?broadcaster_id=141981764
	query := fmt.Sprintf("chat/chatters?%s", values.Encode())
	err = c.doRequest(ctx, "GET", query, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
