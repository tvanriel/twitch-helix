package twitchhelix

import (
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	// httpClient is used to make HTTP requests.
	httpClient HTTPClient

	// baseURL is the base Twitch API URL, typically
	// https://api.twitch.tv/helix/.
	baseURL string

	// clientID is the Twitch application client ID.
	clientID *string

	// token is the OAuth access token associated with the client ID.
	token *string
}

// NewClient initializes and returns a new Twitch API client.
//
// clientID is the Twitch application client ID.
// token is the OAuth access token associated with the client ID.
// httpClient is the HTTP client to use. If nil, http.DefaultClient is used.
//
// To generate an OAuth token with the required scopes, use the Twitch CLI.
// See the Twitch CLI documentation for details.
// Example:
// twitch token -u -s 'analytics:read:extensions analytics:read:games bits:read channel:bot channel:manage:ads channel:read:ads channel:manage:broadcast channel:read:charity channel:manage:clips channel:edit:commercial channel:read:editors channel:manage:extensions channel:read:goals channel:read:guest_star channel:manage:guest_star channel:read:hype_train channel:manage:moderators channel:read:polls channel:manage:polls channel:read:predictions channel:manage:predictions channel:manage:raids channel:read:redemptions channel:manage:redemptions channel:manage:schedule channel:read:stream_key channel:read:subscriptions channel:manage:videos channel:read:vips channel:manage:vips channel:moderate clips:edit editor:manage:clips moderation:read moderator:manage:announcements moderator:manage:automod moderator:read:automod_settings moderator:manage:automod_settings moderator:read:banned_users moderator:manage:banned_users moderator:read:blocked_terms moderator:read:chat_messages moderator:manage:blocked_terms moderator:manage:chat_messages moderator:read:chat_settings moderator:manage:chat_settings moderator:read:chatters moderator:read:followers moderator:read:guest_star moderator:manage:guest_star moderator:read:moderators moderator:read:shield_mode moderator:manage:shield_mode moderator:read:shoutouts moderator:manage:shoutouts moderator:read:suspicious_users moderator:read:unban_requests moderator:manage:unban_requests moderator:read:vips moderator:read:warnings moderator:manage:warnings user:bot user:edit user:edit:broadcast user:read:blocked_users user:manage:blocked_users user:read:broadcast user:read:chat user:manage:chat_color user:read:email user:read:emotes user:read:follows user:read:moderated_channels user:read:subscriptions user:read:whispers user:manage:whispers user:write:chat chat:read chat:edit'
func NewClient(clientID, token *string, httpClient HTTPClient) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{
		httpClient: httpClient,
		baseURL:    "https://api.twitch.tv/helix/",
		clientID:   clientID,
		token:      token,
	}
}
