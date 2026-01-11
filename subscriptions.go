package twitchhelix

import (
	"context"
	"fmt"

	"github.com/google/go-querystring/query"
)

// CheckSubscriptionRequest represents the data required to request if
// a user is subscribed or not to a broadcaster.
type CheckSubscriptionRequest struct {
	// BroadcasterID is the id of the Broadcaster you are checking.
	//
	// Required
	BroadcasterID string `url:"broadcaster_id" json:"broadcaster_id"`

	// UserID is the unique identifier of the user you want to check.
	//
	// Required
	UserID string `url:"user_id" json:"user_id"`
}

// CheckSubscriptionResponse represents the response returned
type CheckSubscriptionResponse struct {
	// Data is a slice of SubscriptionData
	Data []SubscriptionData `json:"data"`
}

// SubscriptionData represents the data about a users subscription status to a broadcaster
type SubscriptionData struct {
	// BroadcasterID is the id of the Broadcaster you are checking.
	BroadcasterID string `json:"broadcaster_id"`

	// BroadcasterLogin is the login name of the broadcaster the user is subscribed to
	BroadcasterLogin string `json:"broadcaster_login"`

	// BroadcasterName is the username of the broadcaster the user is subscribed to
	BroadcasterName string `json:"broadcaster_name"`

	// Tier of the current subscription
	//
	// Tier 1 = "1000"
	// Tier 2 = "2000"
	// Tier 3 = "3000"
	Tier string `json:"tier"`

	// Recieved the subscription as a gift from another user
	IsGift bool `json:"is_gift"`

	// GifterID is the id of the user wo gifted the sub IF it was gifted
	//
	// Possibly Empty
	GifterID *string `json:"gifter_id,omitempty"`

	// GifterLogin is the login of the user wo gifted the sub IF it was gifted
	//
	// Possibly Empty
	GifterLogin *string `json:"gifter_login,omitempty"`

	// GifterName is the username of the user wo gifted the sub IF it was gifted
	//
	// Possibly Empty
	GifterName *string `json:"gifter_name,omitempty"`
}

// CheckSubscription gets the subscription data for a user to a broadcaster
func (c *Client) CheckSubscription(ctx context.Context, req CheckSubscriptionRequest) (*CheckSubscriptionResponse, error) {
	var resp CheckSubscriptionResponse
	// Example Request URL:
	// https://api.twitch.tv/helix/subscriptions/user?broadcaster_id=149747285&user_id=141981764
	query := fmt.Sprintf("subscriptions/user?broadcaster_id=%s&user_id=%s", req.BroadcasterID, req.UserID)
	err := c.doRequest(ctx, "GET", query, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSubscriptionRequest represents the query parameters for fetching broadcaster subscriptions.
type GetSubscriptionsRequest struct {
	// BroadcasterID represents the broadcasters id that is being checked
	BroadcasterID string `url:"broadcaster_id" json:"broadcaster_id"`

	// UserID lets you check for specific users
	// You can only specify up to 100 users
	// If left empty you will get a list of all users
	UserID []string `url:"user_id,omitempty" json:"user_id,omitempty"`

	// First represents how many users you want on each "page"
	//
	// Default 20
	// Max 100
	// Optional
	First *string `url:"first,omitempty" json:"first,omitempty"`

	// After represents the pagination's cursor
	// User after when you want to see the next page
	//
	// Optional
	After *string `url:"after,omitempty" json:"after,omitempty"`

	// Before represents the pagination's cursor
	// User before when you want to see the previous page
	//
	// Optional
	Before *string `url:"before,omitempty" json:"before,omitempty"`
}

// GetSubscriptionResponse represents the response from the Get Broadcaster Subscriptions
//
// It contains a slice of SubscriptionDataV2 and pagination info
type GetSubscriptionsResponse struct {
	// Data contains all subscribers data
	//
	// May be empty if there are no subscribers
	Data []SubscriptionDataV2 `json:"data"`

	// Pagination is the pagination information
	//
	// May be empty if there are no more pages
	Pagination *Pagination `json:"pagination,omitempty"`

	// Points represent the sub points a broadcaster has
	Points int `json:"points"`

	// Total represents the total number of subscribers a broadcaster has
	Total int `json:"total"`
}

// SubscriptionDataV2 represents the data about a users subscription status to a broadcaster
// from a response for GetSubscriptions
type SubscriptionDataV2 struct {
	// BroadcasterID is the id of the Broadcaster you are checking.
	BroadcasterID string `json:"broadcaster_id"`

	// BroadcasterLogin is the login name of the broadcaster the user is subscribed to
	BroadcasterLogin string `json:"broadcaster_login"`

	// BroadcasterName is the username of the broadcaster the user is subscribed to
	BroadcasterName string `json:"broadcaster_name"`

	// UserID is the id of the subscribed user
	UserID string `json:"user_id"`

	// UserName is the login name of the subscribed user
	UserName string `json:"user_name"`

	// UserLogin is the username of the subscribed user
	UserLogin string `json:"user_login"`

	// Tier of the current subscription
	//
	// Tier 1 = "1000"
	// Tier 2 = "2000"
	// Tier 3 = "3000"
	Tier string `json:"tier"`

	// Plan name is the name of the subscription
	PlanName string `json:"plan_name"`

	// Recieved the subscription as a gift from another user
	IsGift bool `json:"is_gift"`

	// GifterID is the id of the user wo gifted the sub IF it was gifted
	//
	// Possibly Empty
	GifterID *string `json:"gifter_id,omitempty"`

	// GifterLogin is the login of the user wo gifted the sub IF it was gifted
	//
	// Possibly Empty
	GifterLogin *string `json:"gifter_login,omitempty"`

	// GifterName is the username of the user wo gifted the sub IF it was gifted
	//
	// Possibly Empty
	GifterName *string `json:"gifter_name,omitempty"`
}

// GetSubscriptions will return the information for subscribers to a broadcaster
//
// The returned slice can be empty if there are no subscribers
func (c *Client) GetSubscriptions(ctx context.Context, req GetSubscriptionsRequest) (*GetSubscriptionsResponse, error) {
	var resp GetSubscriptionsResponse
	values, err := query.Values(req)
	if err != nil {
		return nil, err
	}
	// Example Request URL:
	// https://api.twitch.tv/helix/subscriptions?broadcaster_id=141981764
	query := fmt.Sprintf("subscriptions?%s", values.Encode())
	err = c.doRequest(ctx, "GET", query, nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
