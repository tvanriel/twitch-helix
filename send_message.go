package twitchhelix

import (
	"context"
)

// SendMessageRequest represents the data needed to send a message to a channel
type SendMessageRequest struct {
	// BroadcasterID represents the ID of the broadcaster whose chat you are sending a message to.
	BroadcasterID string `json:"broadcaster_id"`

	// SenderID is the ID of the user that matches the client id.
	SenderID string `json:"sender_id"`

	// Message represents the actual message being sent to the chat.
	Message string `json:"message"`

	// ReplyParentMessageID represents the id of a message that is being replied to.
	//
	// Optional
	ReplyParentMessageID *string `json:"reply_parent_message_id"`

	// ForSourceOnly allows a user to send a message only to host when using shared chat.
	//
	// Optional
	ForSourceOnly *bool `json:"for_source_only"`
}

// SendMessageResponse represents the data recieved from the server when sending a message.
type SendMessageResponse struct {
	// MessageID represents the id of the message sent.
	MessageID string `json:"message_id"`

	// IsSent represents if the message was sent or not.
	IsSent bool `json:"is_sent"`

	// DropReason represents the reason a message was NOT sent.
	//
	// Can be empty
	DropReason *DropReason `json:"drop_reason,omitempty"`
}

// DropReason represents the information for why a message was dropped
type DropReason struct {
	// Code is the error code for a message being dropped.
	Code string `json:"code"`

	// Message is a plane english reason the message was dropped.
	Message string `json:"message"`
}

// SendMessage sends a message to a specified twitch channel.
func (c *Client) SendMessage(ctx context.Context, req SendMessageRequest) (*SendMessageResponse, error) {
	var resp SendMessageResponse
	err := c.doRequest(ctx, "POST", "chat/messages", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
