package twitchhelix

import (
	"context"
	"fmt"

	"github.com/google/go-querystring/query"
)

// RequestCustomReward represents the data used to create or update a custom channel point reward.
type RequestCustomReward struct {
	// Title is the reward's title.
	Title *string `json:"title,omitempty"`

	// Prompt is the description shown to users when redeeming the reward.
	Prompt *string `json:"prompt,omitempty"`

	// Cost is the number of channel points required to redeem the reward.
	Cost *int64 `json:"cost,omitempty"`

	// BackgroundColor is the background color of the reward.
	BackgroundColor *string `json:"background_color,omitempty"`

	// IsEnabled determines whether the reward is enabled.
	IsEnabled *bool `json:"is_enabled,omitempty"`

	// IsUserInputRequired determines whether the user must provide input when redeeming.
	IsUserInputRequired *bool `json:"is_user_input_required,omitempty"`

	// IsMaxPerStreamEnabled determines whether the reward has a per-stream limit.
	IsMaxPerStreamEnabled *bool `json:"is_max_per_stream_enabled,omitempty"`

	// MaxPerStream is the maximum number of redemptions allowed per stream.
	MaxPerStream *int64 `json:"max_per_stream,omitempty"`

	// IsMaxPerUserPerStreamEnabled determines whether the reward has a per-user per-stream limit.
	IsMaxPerUserPerStreamEnabled *bool `json:"is_max_per_user_per_stream_enabled,omitempty"`

	// MaxPerUserPerStream is the maximum number of redemptions allowed per user per stream.
	MaxPerUserPerStream *int64 `json:"max_per_user_per_stream,omitempty"`

	// IsGlobalCooldownEnabled determines whether a global cooldown is enabled.
	IsGlobalCooldownEnabled *bool `json:"is_global_cooldown_enabled,omitempty"`

	// GlobalCooldownSeconds is the cooldown period in seconds between redemptions.
	GlobalCooldownSeconds *int64 `json:"global_cooldown_seconds,omitempty"`

	// IsPaused determines whether the reward is paused.
	IsPaused *bool `json:"is_paused,omitempty"`

	// ShouldRedemptionsSkipRequestQueue determines whether redemptions skip the request queue.
	ShouldRedemptionsSkipRequestQueue *bool `json:"should_redemptions_skip_request_queue,omitempty"`
}

// ResponseCustomReward represents the response returned for custom reward requests.
type ResponseCustomReward struct {
	Data []Reward `json:"data"`
}

// Reward represents a channel point custom reward.
type Reward struct {
	// BroadcasterID is the ID of the broadcaster that owns the reward.
	BroadcasterID string `json:"broadcaster_id"`

	// BroadcasterLogin is the login name of the broadcaster.
	BroadcasterLogin string `json:"broadcaster_login"`

	// BroadcasterName is the display name of the broadcaster.
	BroadcasterName string `json:"broadcaster_name"`

	// ID is the unique identifier of the reward.
	ID string `json:"id"`

	// Title is the reward's title.
	Title string `json:"title"`

	// Prompt is the description shown to users.
	Prompt string `json:"prompt,omitempty"`

	// Cost is the number of channel points required to redeem the reward.
	Cost int64 `json:"cost"`

	// Image is the custom image set for the reward.
	Image *RewardImage `json:"image,omitempty"`

	// DefaultImage is the default image for the reward.
	DefaultImage RewardImage `json:"default_image"`

	// BackgroundColor is the reward's background color.
	BackgroundColor string `json:"background_color"`

	// IsEnabled determines whether the reward is enabled.
	IsEnabled bool `json:"is_enabled"`

	// IsUserInputRequired determines whether user input is required to redeem the reward.
	IsUserInputRequired bool `json:"is_user_input_required"`

	// MaxPerStreamSetting contains the per-stream redemption settings.
	MaxPerStreamSetting MaxPerStreamSetting `json:"max_per_stream_setting"`

	// MaxPerUserPerStreamSetting contains the per-user per-stream redemption settings.
	MaxPerUserPerStreamSetting MaxPerUserPerStreamSetting `json:"max_per_user_per_stream_setting"`

	// GlobalCooldownSetting contains the global cooldown settings.
	GlobalCooldownSetting GlobalCooldownSetting `json:"global_cooldown_setting"`

	// IsPaused determines whether the reward is paused.
	IsPaused bool `json:"is_paused"`

	// IsInStock determines whether the reward is currently redeemable.
	IsInStock bool `json:"is_in_stock"`

	// ShouldRedemptionsSkipRequestQueue determines whether redemptions skip the request queue.
	ShouldRedemptionsSkipRequestQueue bool `json:"should_redemptions_skip_request_queue"`

	// RedemptionsRedeemedCurrentStream is the number of redemptions during the current stream.
	RedemptionsRedeemedCurrentStream *int `json:"redemptions_redeemed_current_stream,omitempty"`

	// CooldownExpiresAt is the timestamp when the global cooldown expires.
	CooldownExpiresAt *string `json:"cooldown_expires_at,omitempty"`
}

// RewardImage represents the image URLs for a reward.
type RewardImage struct {
	URL1x string `json:"url_1x"`
	URL2x string `json:"url_2x"`
	URL4x string `json:"url_4x"`
}

// MaxPerStreamSetting represents per-stream redemption limits.
type MaxPerStreamSetting struct {
	IsEnabled    bool  `json:"is_enabled"`
	MaxPerStream int64 `json:"max_per_stream"`
}

// MaxPerUserPerStreamSetting represents per-user per-stream redemption limits.
type MaxPerUserPerStreamSetting struct {
	IsEnabled           bool  `json:"is_enabled"`
	MaxPerUserPerStream int64 `json:"max_per_user_per_stream"`
}

// GlobalCooldownSetting represents global cooldown settings for a reward.
type GlobalCooldownSetting struct {
	IsEnabled             bool  `json:"is_enabled"`
	GlobalCooldownSeconds int64 `json:"global_cooldown_seconds"`
}

// CreateCustomReward creates a new custom channel point reward for a broadcaster.
// The broadcaster must have channel points enabled.
func (c *Client) CreateCustomReward(ctx context.Context, req RequestCustomReward, broadcasterID string) (*ResponseCustomReward, error) {
	var resp ResponseCustomReward

	query := "channel_points/custom_rewards?broadcaster_id=" + broadcasterID

	err := c.doRequest(ctx, "POST", query, req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateCustomReward updates the fields of an existing custom channel point reward.
// Only the fields provided in the request will be updated.
func (c *Client) UpdateCustomReward(ctx context.Context, req RequestCustomReward, broadcasterID, rewardID string) (*ResponseCustomReward, error) {
	var resp ResponseCustomReward

	query := fmt.Sprintf("channel_points/custom_rewards?broadcaster_id=%s&id=%s", broadcasterID, rewardID)

	err := c.doRequest(ctx, "PATCH", query, req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// RequestGetCustomRewards represents the parameters used to fetch custom rewards.
type RequestGetCustomRewards struct {
	// BroadcasterID is the ID of the broadcaster that owns the rewards.
	BroadcasterID string `url:"broadcaster_id"`

	// RewardID filters the response to specific reward IDs.
	RewardID []string `url:"id,omitempty"`

	// OnlyManagableRewards determines whether only manageable rewards are returned.
	OnlyManagableRewards bool `url:"only_manageable_rewards"`
}

// GetCustomRewards retrieves a list of custom channel point rewards.
// If reward IDs are provided, only those rewards will be returned.
func (c *Client) GetCustomRewards(ctx context.Context, req RequestGetCustomRewards) (*ResponseCustomReward, error) {
	var resp ResponseCustomReward

	values, err := query.Values(req)
	if err != nil {
		return nil, err
	}

	query := "channel_points/custom_rewards?" + values.Encode()

	err = c.doRequest(ctx, "GET", query, nil, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
