package twitchhelix

import (
	"context"

	"github.com/google/go-querystring/query"
)

// RequestSendShoutout represents the data required to send a shoutout to a channel.
type RequestSendShoutout struct {
	// FromBroadcasterID represents the id of the broadcaster that is giving the shoutout.
	FromBroadcasterID string `url:"from_broadcaster_id"`

	// ToBroadcasterID represents the id of the user that is being shouted out.
	ToBroadcasterID string `url:"to_broadcaster_id"`

	// ModeratorID is the id of the user sending the command.
	//
	// This must be the id linked to the client id.
	ModeratorID string `url:"moderator_id"`
}

// SendShoutout sends a shoutout to a user in a channel.
// You must be at least a moderator to perform this action.
// A "/shoutout" can only be sent every 2 minutes. You will receive an error if you query too soon.
func (c *Client) SendShoutout(ctx context.Context, req RequestSendShoutout) error {
	values, err := query.Values(req)
	if err != nil {
		return err
	}

	endpoint := "chat/shoutouts?" + values.Encode()

	err = c.doRequest(ctx, "POST", endpoint, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
