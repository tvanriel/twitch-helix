package twitchhelix

import (
	"context"
)

// RequestStartRaid represents the data required to send a request to start a raid.
type RequestStartRaid struct {
	// FromBroadcasterID represents the id of the broadcaster who is raiding.
	//
	// This id must be linked to the client id.
	FromBroadcasterID string `json:"from_broadcaster_id"`

	// ToBroadcasterID represents the user that is being raided.
	ToBroadcasterID string `json:"to_broadcaster_id"`
}

// ResponseStartRaid represents the data recieved from the server when starting a raid.
type ResponseStartRaid struct {
	// StartRaidData represents the data recieved when starting a raid.
	StartRaidData []StartRaidData `json:"data"`
}

// StartRaidData represents the data about a read starting.
type StartRaidData struct {
	// CreatedAt represents the time a raid was created.
	//
	// A raid will go through 1 minute after IF it is not canceled.
	CreatedAt string `json:"created_at"`

	// IsMature represents if there is a mature 18+ tag on the channel being raided.
	IsMature bool `json:"is_mature"`
}

// StartRaid attempts to start a raid to another channel.
func (c *Client) StartRaid(ctx context.Context, req RequestStartRaid) (*ResponseStartRaid, error) {
	var resp ResponseStartRaid
	err := c.doRequest(ctx, "POST", "raids", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
