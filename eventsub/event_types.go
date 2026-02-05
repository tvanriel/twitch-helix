package twitcheventsub

// Reward represents a channel points reward
type Reward struct {
	// ID is the unique id of the reward
	ID string `json:"id"`

	// Title is the name of the reward
	Title string `json:"title"`

	// Cost is the number of points required to redeem
	Cost int `json:"cost"`

	// Prompt is the reward description
	Prompt string `json:"prompt"`
}

// ChannelPointsRedemptionEvent is triggered when a viewer redeems a reward
type ChannelPointsRedemptionEvent struct {
	// ID is the unique id for this redemption
	ID string `json:"id"`

	// BroadcasterUserID is the broadcaster receiving the redemption
	BroadcasterUserID string `json:"broadcaster_user_id"`

	// BroadcasterUserName is the display name of the broadcaster
	BroadcasterUserName string `json:"broadcaster_user_name"`

	// UserID is the ID of the user redeeming the reward
	UserID string `json:"user_id"`

	// UserName is the display name of the user redeeming the reward
	UserName string `json:"user_name"`

	// UserInput is an optional message provided by the user
	UserInput string `json:"user_input"`

	// Status indicates redemption state: "FULFILLED", "CANCELED", "UNFULFILLED"
	Status string `json:"status"`

	// RedeemedAt is the timestamp when the reward was redeemed
	RedeemedAt string `json:"redeemed_at"`

	// Reward contains details of the redeemed reward
	Reward Reward `json:"reward"`
}

// ChannelRaidEvent is triggered when a broadcaster raids another channel
type ChannelRaidEvent struct {
	// FromBroadcasterUserID is the ID of the raiding broadcaster
	FromBroadcasterUserID string `json:"from_broadcaster_user_id"`

	// FromBroadcasterUserLogin is the login of the raiding broadcaster
	FromBroadcasterUserLogin string `json:"from_broadcaster_user_login"`

	// FromBroadcasterUserName is the display name of the raiding broadcaster
	FromBroadcasterUserName string `json:"from_broadcaster_user_name"`

	// ToBroadcasterUserID is the ID of the raided broadcaster
	ToBroadcasterUserID string `json:"to_broadcaster_user_id"`

	// ToBroadcasterUserLogin is the login of the raided broadcaster
	ToBroadcasterUserLogin string `json:"to_broadcaster_user_login"`

	// ToBroadcasterUserName is the display name of the raided broadcaster
	ToBroadcasterUserName string `json:"to_broadcaster_user_name"`

	// Viewers is the number of viewers in the raid
	Viewers int `json:"viewers"`
}

// AdBreakEvent is triggered when an ad break starts
type AdBreakEvent struct {
	// DurationSeconds is the length of the ad in seconds
	DurationSeconds int `json:"duration_seconds"`

	// StartedAt is the timestamp when the ad started
	StartedAt string `json:"started_at"`

	// IsAutomatic indicates if the ad was automatically triggered
	IsAutomatic bool `json:"is_automatic"`

	// BroadcasterUserID is the ID of the broadcaster running the ad
	BroadcasterUserID string `json:"broadcaster_user_id"`

	// BroadcasterUserLogin is the login of the broadcaster running the ad
	BroadcasterUserLogin string `json:"broadcaster_user_login"`

	// BroadcasterUserName is the display name of the broadcaster
	BroadcasterUserName string `json:"broadcaster_user_name"`

	// RequesterUserID is the ID of the user who requested the ad (manual)
	RequesterUserID string `json:"requester_user_id"`

	// RequesterUserLogin is the login of the requester
	RequesterUserLogin string `json:"requester_user_login"`

	// RequesterUserName is the display name of the requester
	RequesterUserName string `json:"requester_user_name"`
}

// StreamOnlineEvent is triggered when a broadcaster goes live
type StreamOnlineEvent struct {
	// ID is the unique id of the stream
	ID string `json:"id"`

	// BroadcasterUserID is the broadcaster going live
	BroadcasterUserID string `json:"broadcaster_user_id"`

	// BroadcasterUserLogin is the login of the broadcaster
	BroadcasterUserLogin string `json:"broadcaster_user_login"`

	// BroadcasterUserName is the display name of the broadcaster
	BroadcasterUserName string `json:"broadcaster_user_name"`

	// Type is the stream type (usually "live")
	Type string `json:"type"`

	// StartedAt is the timestamp when the stream started
	StartedAt string `json:"started_at"`
}

// StreamOfflineEvent is triggered when a broadcaster goes offline
type StreamOfflineEvent struct {
	// BroadcasterUserID is the broadcaster going offline
	BroadcasterUserID string `json:"broadcaster_user_id"`

	// BroadcasterUserLogin is the login of the broadcaster
	BroadcasterUserLogin string `json:"broadcaster_user_login"`

	// BroadcasterUserName is the display name of the broadcaster
	BroadcasterUserName string `json:"broadcaster_user_name"`
}

// ChannelUpdateEvent is triggered when a broadcaster updates their channel information
type ChannelUpdateEvent struct {
	// BroadcasterUserID is the ID of the broadcaster
	BroadcasterUserID string `json:"broadcaster_user_id"`

	// BroadcasterUserLogin is the login of the broadcaster
	BroadcasterUserLogin string `json:"broadcaster_user_login"`

	// BroadcasterUserName is the display name of the broadcaster
	BroadcasterUserName string `json:"broadcaster_user_name"`

	// Title is the updated stream title
	Title string `json:"title"`

	// Language is the updated stream language (ISO 639-1)
	Language string `json:"language"`

	// CategoryID is the updated game/category ID
	CategoryID string `json:"category_id"`

	// ContentClassificationLabels are the updated content classification labels
	ContentClassificationLables []string `json:"content_classification_labels"`
}

// ChannelSubscriptionGiftEvent is triggered when a user gifts a subscription
type ChannelSubscriptionGiftEvent struct {
	// UserID is the ID of the user redeeming the reward
	UserID string `json:"user_id"`

	// UserName is the display name of the user redeeming the reward
	UserName string `json:"user_name"`

	// UserInput is an optional message provided by the user
	UserInput string `json:"user_input"`

	// BroadcasterUserID is the ID of the broadcaster
	BroadcasterUserID string `json:"broadcaster_user_id"`

	// BroadcasterUserLogin is the login of the broadcaster
	BroadcasterUserLogin string `json:"broadcaster_user_login"`

	// BroadcasterUserName is the display name of the broadcaster
	BroadcasterUserName string `json:"broadcaster_user_name"`

	// Total Represents the total amount of subs gifted in this single event
	Total int `json:"total"`

	// Tier represents the tier of subscription gifted ("1000", "2000", "3000")
	Tier string `json:"tier"`

	// CumulativeTotal represents to total number of subscriptions that a user has gifted to the Broadvasters channel
	CumulativeTotal int `json:"cumulative_total"`

	// IsAnonymous represents if the user checked to remain anonymous when gifting the subscriptions
	IsAnonymous bool `json:"is_anonymous"`
}

// ChatMessageFragment is the fragment field in the [ChannelChatMessagePayloadBody] struct.
type ChatMessageFragment struct {
	Type string `json:"type"`
	Text string `json:"text"`

	CheerEmote any `json:"cheer_emote"` // TODO: figure out what these types are. Twitch docs are vague.
	Emote      any `json:"emote"`
	Mention    any `json:"mention"`
}

// ChannelChatMessagePayloadBody is the message field in the [ChannelChatMessagePayload] struct.
type ChannelChatMessagePayloadBody struct {
	Text      string                `json:"text"`
	Fragments []ChatMessageFragment `json:"fragments"`
}

type Badge struct {
	SetID string `json:"set_id"`
	ID    string `json:"id"`
	Info  string `json:"info"`
}

// ChannelChatMessagePayload is the payload received from Twitch on the `channel.chat.message` notification type.
type ChannelChatMessagePayload struct {
	// BroadcasterUserID is the userid of the broadcaster whose channel the message was sent to.
	BroadcasterUserID string `json:"broadcaster_user_id"`

	// BroadcasterUserLogin is the login of the broadcaster whose channel the message was sent to.
	BroadcasterUserLogin string `json:"broadcaster_user_login"`

	// BroadcasterUserUserName is the username of the broadcaster whose channel the message was sent to.
	BroadcasterUserUserName string `json:"broadcaster_user_user_name"`

	// ChatterUserID is the userid of the chat message author.
	ChatterUserID string `json:"chatter_user_id"`

	// ChatterUserLogin is the login of the chat message author.
	ChatterUserLogin string `json:"chatter_user_login"`

	// ChatterUserName is the username of the chat message author.
	ChatterUserName string `json:"chatter_user_name"`

	// MessageID is the uuid of this message.
	MessageID string `json:"messageId"`

	// Message is the display info of this chat message.
	Message ChannelChatMessagePayloadBody `json:"message"`

	// Color is the authors username colour.
	Color string `json:"color"`

	// Badges is the list of badges that should be rendered.
	Badges []Badge `json:"badges"`

	// Cheer TODO figure out what this is.
	Cheer any `json:"cheer"`

	// Cheer TODO figure out what this is.
	Reply any `json:"reply"`

	// Cheer TODO figure out what this is.
	ChannelPointsCustomRewardID any `json:"channel_points_custom_reward_id"`

	// SourceBroadcasterUserID is the UserID of the source broadcaster, if the chat message was sent in a shared chatbox.
	SourceBroadcasterUserID *string `json:"source_broadcasterUserId"`

	// SourceBroadcasterUserID is the Login of the source broadcaster, if the chat message was sent in a shared chatbox.
	SourceBroadcasterUserLogin *string `json:"source_broadcaster_user_login"`

	// SourceBroadcasterUserID is the UserName of the source broadcaster, if the chat message was sent in a shared chatbox.
	SourceBroadcasterUserName *string `json:"source_broadcaster_user_name"`

	// SourceMessageID is the ID of the original message , if the chat message was sent in a shared chatbox.
	SourceMessageID *string `json:"source_message_id"`

	// SourceBadges is the badges to render if the chat message was sent in a shared chatbox.
	SourceBadges []Badge `json:"source_badges"`

	// IsSourceOnly is set if the message is not to be sent to the other channels in a shared chatbox.
	IsSourceOnly bool `json:"is_source_only"`
}
