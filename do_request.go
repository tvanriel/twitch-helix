package twitchhelix

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
)

// doRequest performs an HTTP request and decodes the response.
//
// ctx controls cancellation and timeouts.
// method is the HTTP method (GET, POST, etc.).
// endpoint is the API path without the base URL.
// body is encoded as JSON and sent as the request body.
// out is decoded from the JSON response body.
func (c *Client) doRequest(ctx context.Context, method string, endpoint string, body interface{}, out interface{}) error {
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+endpoint, bodyReader)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+*c.token)
	req.Header.Set("Client-Id", *c.clientID)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return &TwitchAPIError{
			StatusCode: resp.StatusCode,
			Body:       bodyBytes,
		}
	}

	if out != nil {
		return json.NewDecoder(resp.Body).Decode(out)
	}
	return nil
}
