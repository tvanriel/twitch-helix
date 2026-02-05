package twitchhelix

import (
	"context"

	"github.com/google/go-querystring/query"
)

// GetUsersRequest represents the query parameters used to fetch users.
type GetUsersRequest struct {
	// ID filters the request by user IDs.
	//
	// You may specify up to 100 user IDs.
	ID []string `url:"id,omitempty"`

	// Login filters the request by user login names.
	//
	// You may specify up to 100 login names.
	Login []string `url:"login,omitempty"`
}

// GetUsersResponse represents the response returned from the Get Users endpoint.
type GetUsersResponse struct {
	// Data contains the list of users.
	Data []User `json:"data"`
}

// User represents a Twitch user.
type User struct {
	// ID is the unique identifier of the user.
	ID string `json:"id"`

	// Login is the user's login name.
	Login string `json:"login"`

	// DisplayName is the user's display name.
	DisplayName string `json:"display_name"`

	// Type describes the type of user account.
	Type string `json:"type"`

	// BroadcasterType describes the broadcaster status of the user.
	BroadcasterType string `json:"broadcaster_type"`

	// Description is the user's profile description.
	Description string `json:"description"`

	// ProfileImageURL is the URL of the user's profile image.
	ProfileImageURL string `json:"profile_image_url"`

	// OfflineImageURL is the URL of the user's offline image.
	OfflineImageURL string `json:"offline_image_url"`

	// ViewCount is the total number of views for the user.
	// This is CCV NOT users connected to chat.
	//
	// Deprecated.
	ViewCount int64 `json:"view_count"`

	// Email is the user's email address.
	//
	// This field is only returned if the user authorized with the user:read:email scope.
	Email string `json:"email"`

	// CreatedAt is the timestamp when the user account was created.
	//
	// The timestamp is in RFC3339 format.
	CreatedAt string `json:"created_at"`
}

// GetUsers retrieves information about one or more Twitch users.
// Users can be looked up by ID, login name, or both.
func (c *Client) GetUsers(ctx context.Context, req GetUsersRequest) (*GetUsersResponse, error) {
	var resp GetUsersResponse

	values, err := query.Values(req)
	if err != nil {
		return nil, err
	}

	query := "users?" + values.Encode()

	err = c.doRequest(ctx, "GET", query, nil, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
