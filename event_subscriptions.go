package twitchhelix

import (
	"context"
	"time"
)

// EventRequest represents a request to create an EventSub subscription.
type EventRequest struct {
	// Type is the EventSub subscription type.
	Type string `json:"type"`

	// Version is the version of the EventSub subscription.
	Version string `json:"version"`

	// Conditoin contains the condition required for the subscription.
	Conditoin any `json:"condition"`

	// Transport defines how events are delivered.
	Transport WebsocketTransport `json:"transport"`
}

// WebsocketTransport represents a WebSocket transport configuration.
type WebsocketTransport struct {
	// Method is the transport method used for delivery.
	Method string `json:"method"`

	// SessionID is the WebSocket session identifier.
	SessionID string `json:"session_id"`

	// ConnectedAt is the timestamp when the WebSocket connected.
	//
	// Response only.
	ConnectedAt *time.Time `json:"connected_at"`

	// DisconnectedAt is the timestamp when the WebSocket disconnected.
	//
	// Response only.
	DisconnectedAt *time.Time `json:"disconnected_at"`
}

// =============================================================

// ConditionStreamOnline represents the condition for a stream going online event.
type ConditionStreamOnline struct {
	// BroadcasterUserID is the ID of the broadcaster to monitor.
	BroadcasterUserID string `json:"broadcaster_user_id"`
}

// EventStreamOnline subscribes to stream.online events for a broadcaster.
func (c *Client) EventStreamOnline(ctx context.Context, sessionID string, condition ConditionStreamOnline) (*any, error) {
	req := EventRequest{
		Type:      "stream.online",
		Version:   "1",
		Conditoin: condition,
		Transport: WebsocketTransport{
			Method:    "websocket",
			SessionID: sessionID,
		},
	}
	var resp any
	err := c.doRequest(ctx, "POST", "eventsub/subscriptions", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ConditionStreamOffline represents the condition for a stream going offline event.
type ConditionStreamOffline struct {
	// BroadcasterUserID is the ID of the broadcaster to monitor.
	BroadcasterUserID string `json:"broadcaster_user_id"`
}

// EventStreamOffline subscribes to stream.offline events for a broadcaster.
func (c *Client) EventStreamOffline(ctx context.Context, sessionID string, condition ConditionStreamOffline) (*any, error) {
	req := EventRequest{
		Type:      "stream.offline",
		Version:   "1",
		Conditoin: condition,
		Transport: WebsocketTransport{
			Method:    "websocket",
			SessionID: sessionID,
		},
	}
	var resp any
	err := c.doRequest(ctx, "POST", "eventsub/subscriptions", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ConditionChannelUpdate represents the condition for a channel update event.
type ConditionChannelUpdate struct {
	// BroadcasterUserID is the ID of the broadcaster to monitor.
	BroadcasterUserID string `json:"broadcaster_user_id"`
}

// EventChannelUpdate subscribes to channel.update events for a broadcaster.
func (c *Client) EventChannelUpdate(ctx context.Context, sessionID string, condition ConditionChannelUpdate) (*any, error) {
	req := EventRequest{
		Type:      "channel.update",
		Version:   "2",
		Conditoin: condition,
		Transport: WebsocketTransport{
			Method:    "websocket",
			SessionID: sessionID,
		},
	}
	var resp any
	err := c.doRequest(ctx, "POST", "eventsub/subscriptions", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ConditionChannelRaid represents the condition for a channel raid event.
//
// You must specify only one broadcaster ID.
type ConditionChannelRaid struct {
	// FromBroadcasterUserID is the ID of the broadcaster sending the raid.
	FromBroadcasterUserID *string `json:"from_broadcaster_user_id"`

	// ToBroadcasterUserID is the ID of the broadcaster receiving the raid.
	ToBroadcasterUserID *string `json:"to_broadcaster_user_id"`
}

// EventChannelRaid subscribes to channel.raid events.
func (c *Client) EventChannelRaid(ctx context.Context, sessionID string, condition ConditionChannelRaid) (*any, error) {
	req := EventRequest{
		Type:      "channel.raid",
		Version:   "1",
		Conditoin: condition,
		Transport: WebsocketTransport{
			Method:    "websocket",
			SessionID: sessionID,
		},
	}
	var resp any
	err := c.doRequest(ctx, "POST", "eventsub/subscriptions", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ConditionEventChannelPointsCustomRewardRedemptionAdd represents the condition for a reward redemption event.
type ConditionEventChannelPointsCustomRewardRedemptionAdd struct {
	// BroadcasterUserID is the ID of the broadcaster to monitor.
	BroadcasterUserID string `json:"broadcaster_user_id"`
}

// EventChannelPointsCustomRewardRedemptionAdd subscribes to channel point reward redemption events.
func (c *Client) EventChannelPointsCustomRewardRedemptionAdd(ctx context.Context, sessionID string, condition ConditionEventChannelPointsCustomRewardRedemptionAdd) (*any, error) {
	req := EventRequest{
		Type:      "channel.channel_points_custom_reward_redemption.add",
		Version:   "1",
		Conditoin: condition,
		Transport: WebsocketTransport{
			Method:    "websocket",
			SessionID: sessionID,
		},
	}
	var resp any
	err := c.doRequest(ctx, "POST", "eventsub/subscriptions", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ConditionChannelAdBreakBegin represents the condition for an ad break begin event.
type ConditionChannelAdBreakBegin struct {
	// BroadcasterUserID is the ID of the broadcaster.
	BroadcasterUserID *string `json:"broadcaster_user_id"`
}

// EventChannelAdBreakBegin subscribes to channel.ad_break.begin events.
func (c *Client) EventChannelAdBreakBegin(ctx context.Context, sessionID string, condition ConditionChannelAdBreakBegin) (*any, error) {
	req := EventRequest{
		Type:      "channel.ad_break.begin",
		Version:   "1",
		Conditoin: condition,
		Transport: WebsocketTransport{
			Method:    "websocket",
			SessionID: sessionID,
		},
	}
	var resp any
	err := c.doRequest(ctx, "POST", "eventsub/subscriptions", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
