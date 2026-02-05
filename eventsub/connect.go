package twitcheventsub

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// SessionConfig represents the basic information to establish a connection
type SessionConfig struct {
	// Session is a pointer that represents the sessions data
	Session *Session

	// Events represents the channel that events will be sent to
	Events chan<- Event

	// mu represents a RWMutex to pretect session data
	mu sync.RWMutex

	logger *zap.Logger
}

// Event represents the information of a received event
type Event struct {
	// MessageType represents the type of message
	// Welcome, Keepalive, Notification, Reconnect, Revocation
	MessageType string

	// SubscriptionType represents the event type
	SubscriptionType string

	// Data represents all data for the event
	Data []byte
}

// TwitchEventSubURL is the default URL for the websocket server where twitch hosts EventSub
var TwitchEventSubURL = "wss://eventsub.wss.twitch.tv/ws"

// NewEventSubWebsocket is a init function for session config
func NewEventSubWebsocket(eventMessageChan chan<- Event) *SessionConfig {
	s := &SessionConfig{
		Events: eventMessageChan,
		Session: &Session{
			ReconnectURL: &TwitchEventSubURL,
		},
		logger: zap.L(),
	}

	return s
}

func (s *SessionConfig) SetLogger(logger *zap.Logger) {
	s.logger = logger
}

// Connect opens a connection to Twitch Event Sub
func (s *SessionConfig) Connect() error {
	// Dial the Twitch WebSocket
	c, _, err := websocket.DefaultDialer.Dial(*s.Session.ReconnectURL, nil)
	if err != nil {
		return fmt.Errorf("twitch-helix dial websocket: %w", err)
	}
	defer c.Close()

	// Read messages from Twitch
	for {
		websocketMessageType, message, err := c.ReadMessage()
		if err != nil {
			if websocketMessageType == websocket.PingMessage {
				s.logger.Debug("Websocket PING, but read message has error", zap.Error(err))
				c.WriteMessage(websocket.PongMessage, nil)

				continue
			}

			s.logger.Error("Read Error:", zap.Error(err))

			return nil
		}

		if websocketMessageType == websocket.PingMessage {
			s.logger.Debug("Websocket PING")
			c.WriteMessage(websocket.PongMessage, nil)

			continue
		}

		var msg WebsocketMessage

		err = json.Unmarshal(message, &msg)
		if err != nil {
			s.logger.Warn("twitch-helix decode message", zap.Error(err))

			continue
		}

		logger := s.logger.With(zap.String("message_id", msg.Metadata.MessageID))

		// Handle the different message types
		switch msg.Metadata.MessageType {
		case "session_welcome":
			var parsedMessage WelcomeMessage

			err := json.Unmarshal(message, &parsedMessage)
			if err != nil {
				logger.Error("decode session_welcome essage", zap.Error(err))

				continue
			}

			s.mu.Lock()
			s.Session = &parsedMessage.Payload.Session
			s.mu.Unlock()

			s.Events <- Event{MessageType: parsedMessage.Metadata.MessageType, SubscriptionType: parsedMessage.Metadata.SubscriptionType, Data: message}

		case "session_keepalive":
			var parsedMessage KeepaliveMessage

			err := json.Unmarshal(message, &parsedMessage)
			if err != nil {
				logger.Error("decode session_eepalive message", zap.Error(err))

				continue
			}

			s.Events <- Event{MessageType: parsedMessage.Metadata.MessageType, SubscriptionType: parsedMessage.Metadata.SubscriptionType, Data: message}

		case "notification":
			var parsedMessage NotificationMessage

			err := json.Unmarshal(message, &parsedMessage)
			if err != nil {
				logger.Error("decode notification message", zap.Error(err))

				continue
			}

			s.Events <- Event{MessageType: parsedMessage.Metadata.MessageType, SubscriptionType: parsedMessage.Metadata.SubscriptionType, Data: message}

		case "reconnect_message":
			var parsedMessage ReconnectMessage

			err := json.Unmarshal(message, &parsedMessage)
			if err != nil {
				logger.Error("decode reconnect_message message", zap.Error(err))

				continue
			}

			s.mu.Lock()
			logger.Info("reconnect url updated", zap.Stringp("new_url", parsedMessage.Payload.Session.ReconnectURL))
			s.Session.ReconnectURL = parsedMessage.Payload.Session.ReconnectURL
			s.mu.Unlock()

			if err := c.Close(); err != nil {
				logger.Error("failed to close connection afer twitch asked for reconnect", zap.Error(err))
			}

			if err := s.Connect(); err != nil {
				logger.Error("failed to reconnect after twitch asked for reconnect", zap.Error(err))
				close(s.Events)

				return err
			}
		}
	}
}
