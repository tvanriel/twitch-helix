package twitchhelix

import (
	"context"
	"fmt"
)

// RequestModifyChannelInformation represents the fields that can be updated for a channel.
// At least one field must be included when sending a request.
type RequestModifyChannelInformation struct {
	// GameID is the ID of the game currently being played on the channel.
	GameID *string `json:"game_id,omitempty"`

	// BroadcasterLanguage is the broadcaster's primary language.
	//
	// The value must be an ISO 639-1 two-letter language code.
	BroadcasterLanguage *string `json:"broadcaster_language,omitempty"`

	// Title is the new stream title for the channel.
	Title *string `json:"title,omitempty"`

	// Delay specifies the number of seconds to delay the live stream.
	//
	// Optional, may be 0.
	Delay *int `json:"delay,omitempty"`

	// Tags is a list of tag IDs to associate with the channel.
	Tags *[]string `json:"tags,omitempty"`

	// ContentClassificationLabels specifies content classification labels for the channel.
	ContentClassificationLabels *[]ContentClassificationLabel `json:"content_classification_labels,omitempty"`

	// IsBrandedContent indicates whether the channel contains branded content.
	IsBrandedContent *bool `json:"is_branded_content,omitempty"`
}

// ContentClassificationLabel represents a content classification label and its enabled state.
type ContentClassificationLabel struct {
	// ID is the unique identifier of the label.
	ID string `json:"id"`

	// IsEnabled determines whether the label is applied to the channel.
	IsEnabled bool `json:"is_enabled"`
}

// ModifyChannelInformation updates a broadcaster's channel settings.
// At least one field in RequestModifyChannelInformation must be included.
// Requires broadcaster authentication.
func (c *Client) ModifyChannelInformation(ctx context.Context, req RequestModifyChannelInformation, broadcasterID string) error {
	query := fmt.Sprintf("channels?broadcaster_id=%s", broadcasterID)
	err := c.doRequest(ctx, "PATCH", query, req, nil)
	if err != nil {
		return err
	}
	return nil
}
