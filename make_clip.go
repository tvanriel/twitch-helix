package twitchhelix

import (
	"context"
)

// MakeClipResponse represents the data received after creating a clip
type MakeClipResponse struct {
	// ID represents the id of the clip
	ID string `json:"id"`

	// EditURL represents the url that will allow the creator (you) to edit the url
	EditURL string `json:"edit_url"`
}

// MakeClip creates a clip at the time called
// Twitch tries to capture the previous 90 seconds and provide the edit url
// Twitch will only publish the last 30s of the clip by default
//
// If nothing is returned after 15 seconds, assume clip creation has failed
func (c *Client) MakeClip(ctx context.Context, broadcaster_id string) (*MakeClipResponse, error) {
	var resp MakeClipResponse

	err := c.doRequest(ctx, "POST", "clips?broadcaster_id="+broadcaster_id, nil, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
