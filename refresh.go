package twitchhelix

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type NewToken struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	ExpiresIn    int      `json:"expires_in"`
	Scope        []string `json:"scope"`
	TokenType    string   `json:"token_type"`
}

const twitchRefreshLink = "https://id.twitch.tv/oauth2/token"

func (c *Client) Refresh(clientSecret, refreshToken string) (*NewToken, error) {
	data := url.Values{}
	data.Add("client_id", *c.clientID)
	data.Add("client_secret", clientSecret)
	data.Add("grant_type", "refresh_token")
	data.Add("refresh_token", refreshToken)

	tokenRequestData := data.Encode()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, twitchRefreshLink, strings.NewReader(tokenRequestData))
	if err != nil {
		return nil, fmt.Errorf("failed to create new http request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	var tokenData NewToken

	err = json.Unmarshal(body, &tokenData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.token = &tokenData.AccessToken

	return &tokenData, nil
}

type NewAppToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func (c *Client) RefreshApp(clientSecret string) (*NewAppToken, error) {
	data := url.Values{}
	data.Add("client_id", *c.clientID)
	data.Add("client_secret", clientSecret)
	data.Add("grant_type", "client_credentials")

	tokenRequestData := data.Encode()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, twitchRefreshLink, strings.NewReader(tokenRequestData))
	if err != nil {
		return nil, fmt.Errorf("failed to create new http request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %w", err)
	}

	var tokenData NewAppToken

	err = json.Unmarshal(body, &tokenData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	c.token = &tokenData.AccessToken

	return &tokenData, nil
}
