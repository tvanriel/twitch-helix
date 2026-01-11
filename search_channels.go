package twitchhelix

import (
	"context"
	"fmt"
	"github.com/google/go-querystring/query"
)

// RequestSearchChannels represents the query parameters used to search channels.
type RequestSearchChannels struct {
	// Query is the search string used to find channels.
	//
	// You can search a game name, username, anything you would normally put in the search bar.
	Query string `url:"query" validate:"required"`

	// LiveOnly determines whether only live channels are returned.
	//
	// Default: false
	// Optional
	LiveOnly *bool `url:"live_only,omitempty"`

	// First specifies the maximum number of results to return.
	//
	// The value must be between 1 and 100.
	// Default: 20
	First *int `url:"first,omitempty"`

	// After is the cursor used to fetch the next page of results.
	//
	// Optional
	After *string `url:"after,omitempty"`
}

// ResponseSearchChannels represents the response returned from the Search Channels endpoint.
type ResponseSearchChannels struct {
	// Data contains the list of channels that match the search query.
	Data []Channel `json:"data"`

	// Pagination contains the pagination cursor.
	Pagination *Pagination `json:"pagination,omitempty"`
}

// Channel represents a Twitch channel returned in search results.
type Channel struct {
	// BroadcasterLanguage is the broadcaster's language.
	//
	// The value is an ISO 639-1 language code.
	BroadcasterLanguage string `json:"broadcaster_language"`

	// BroadcasterLogin is the broadcaster's login name.
	BroadcasterLogin string `json:"broadcaster_login"`

	// DisplayName is the broadcaster's display name.
	DisplayName string `json:"display_name"`

	// GameID is the ID of the game currently being played.
	GameID string `json:"game_id"`

	// GameName is the name of the game currently being played.
	GameName string `json:"game_name"`

	// ID is the unique identifier of the broadcaster.
	ID string `json:"id"`

	// IsLive indicates whether the broadcaster is currently live.
	IsLive bool `json:"is_live"`

	// Tags is the list of tags applied to the channel.
	Tags []string `json:"tags"`

	// ThumbnailURL is the URL of the channel's thumbnail image.
	ThumbnailURL string `json:"thumbnail_url"`

	// Title is the current stream title.
	Title string `json:"title"`

	// StartedAt is the timestamp when the stream started.
	//
	// The timestamp is in RFC3339 format.
	StartedAt string `json:"started_at"`
}

// SearchChannels searches for channels that match the specified query.
// Results can be filtered to include only live channels.
func (c *Client) SearchChannels(ctx context.Context, req RequestSearchChannels) (*ResponseSearchChannels, error) {
	var resp ResponseSearchChannels
	values, err := query.Values(req)
	if err != nil {
		return nil, err
	}
	endpoint := fmt.Sprintf("search/channels?%s", values.Encode())
	err = c.doRequest(ctx, "GET", endpoint, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
