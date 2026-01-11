package twitchhelix

import "fmt"

// TwitchAPIError represents an error response returned by the Twitch Helix API.
type TwitchAPIError struct {
	// StatusCode is the HTTP status code returned by the API.
	StatusCode int

	// Body contains the raw response body returned by the API.
	Body []byte
}

func (e *TwitchAPIError) Error() string {
	return fmt.Sprintf("Twitch API Error(%d): %s", e.StatusCode, e.Body)
}
