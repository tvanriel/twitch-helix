package twitchhelix

import (
	"context"

	"github.com/google/go-querystring/query"
)

// StreamRequest represents the query parameters used to fetch streams.
type StreamRequest struct {
	// UserID filters streams by broadcaster user IDs.
	//
	// You may specify up to 100 user IDs.
	UserID []*string `json:"user_id,omitempty" url:"user_id,omitempty"`

	// UserLogin filters streams by broadcaster login names.
	//
	// You may specify up to 100 login names.
	UserLogin []*string `json:"user_login,omitempty" url:"user_login,omitempty"`

	// GameID filters streams by game IDs.
	//
	// You may specify up to 100 game IDs.
	GameID []*string `json:"game_id,omitempty" url:"game_id,omitempty"`

	// Type filters streams by stream type.
	//
	// Valid values are "all" or "live".
	Type *string `json:"type,omitempty" url:"type,omitempty"`

	// Language filters streams by broadcast language.
	//
	// The value must be an ISO 639-1 language code.
	Language []*string `json:"language,omitempty" url:"language,omitempty"`

	// First specifies the maximum number of items to return.
	//
	// The value must be between 1 and 100.
	First *int `json:"first,omitempty" url:"first,omitempty"`

	// Before is the cursor used to fetch the previous page.
	Before *string `json:"before,omitempty" url:"before,omitempty"`

	// After is the cursor used to fetch the next page.
	After *string `json:"after,omitempty" url:"after,omitempty"`
}

// StreamResponse represents the response returned from the Get Streams endpoint.
type StreamResponse struct {
	// Data is the list of live streams.
	Data []*StreamData `json:"data,omitempty"`

	// Pagination contains the pagination cursor.
	Pagination Pagination `json:"pagination"`
}

// StreamData represents the data for a single live stream.
type StreamData struct {
	// ID is the unique identifier of the stream.
	ID *string `json:"id,omitempty"`

	// UserID is the ID of the broadcaster.
	UserID *string `json:"user_id,omitempty"`

	// UserLogin is the login name of the broadcaster.
	UserLogin *string `json:"user_login,omitempty"`

	// UserName is the display name of the broadcaster.
	UserName *string `json:"user_name,omitempty"`

	// GameID is the ID of the game being streamed.
	GameID *string `json:"game_id,omitempty"`

	// GameName is the name of the game being streamed.
	GameName *string `json:"game_name,omitempty"`

	// Type is the stream type.
	//
	// This value is "live" for active streams.
	Type *string `json:"type,omitempty"`

	// Title is the stream title.
	Title *string `json:"title,omitempty"`

	// Tags is the list of tags applied to the stream.
	Tags []*string `json:"tags,omitempty"`

	// ViewerCount is the current number of viewers.
	ViewerCount *int `json:"viewer_count,omitempty"`

	// StartedAt is the timestamp when the stream started.
	//
	// The timestamp is in RFC3339 format.
	StartedAt *string `json:"started_at,omitempty"`

	// Language is the broadcast language.
	//
	// The value is an ISO 639-1 language code.
	Language *string `json:"language,omitempty"`

	// ThumbnailURL is the URL template for the stream thumbnail.
	//
	// Replace {width} and {height} with the desired dimensions.
	ThumbnailURL *string `json:"thumbnail_url,omitempty"`

	// TagIDs is a deprecated field and is always empty.
	TagIDs []*string `json:"tag_ids,omitempty"`

	// IsMature indicates whether the stream is intended for mature audiences.
	IsMature *bool `json:"is_mature,omitempty"`
}

// GetStreams retrieves a list of live streams.
// Results can be filtered using the provided query parameters.
func (c *Client) GetStreams(ctx context.Context, req StreamRequest) (*StreamResponse, error) {
	var resp StreamResponse

	values, err := query.Values(req)
	if err != nil {
		return nil, err
	}

	query := "streams?" + values.Encode()

	err = c.doRequest(ctx, "GET", query, nil, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
