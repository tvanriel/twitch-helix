package twitchhelix

// Pagination represents the cursor-based pagination information
// returned by Twitch API endpoints.
//
// The Cursor field is a string used to fetch the next (or previous)
// page of results. It can be empty if there is no more data.
type Pagination struct {
	// Cursor is the pagination token returned by the API.
	// Include this value in your next request as the `after` parameter
	// to fetch the next page of results.
	// If empty, there are no more pages.
	Cursor string `json:"cursor,omitempty"`
}
