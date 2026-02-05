# twitch-eventsub
"twitch-eventsub" is a Go library for interacting with Twitch’s EventSub API, allowing you to listen to events such as follows, subscriptions, donations, and more.

To register EventSub subscriptions, use [`twitch-helix`](https://github.com/v0idzzy/twitch-helix).
EventSub subscriptions cannot be created until Twitch assigns a WebSocket session.
Twitch provides this session ID in the session_welcome message.

For detailed field info, see the [pkg.go.dev documentation](https://pkg.go.dev/github.com/v0idzzy/twitch-eventsub).
# Features
- Creates a connection to Twitch Event Sub over Web Socket
- Manages state of the sessions
- Decodes all events received # Installation 
```bash
go get github.com/v0idzzy/twitch-eventsub
```
# Example
```go
package main

import (
    "log"
    "context"
    "time"

    "github.com/v0idzzy/twitch-eventsub"
    "github.com/v0idzzy/twitch-helix"
)

func main() {
    oauth := "YOUR_TWITCH_OAUTH"       // no OAuth Prefix
    clientID := "YOUR_CLIENT_ID"       // Client ID linked with the oauth

    helixClient := twitchhelix.NewClient(clientID, oauth, nil)

    eventSubChan := make(chan twitchkiteventsub.Event)
    eventSubWebsocket := twitchkiteventsub.NewEventSubWebsocket(eventSubChan)
    
    go func() {
        eventSubWebsocket.Connect()
    }

    firstWelcome := true // Guarding against multiple session welcomes
    for event := range eventSubChan {
        switch event.MessageType {
        case "session_welcome":
            if firstWelcome {
                firstWelcome = false

                //initialze connections
                ctx := context.Background()
                ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
                defer cancel()

                _, err := helixClient.EventStreamOnline(
                    ctx,
                    eventSubWebsocket.Session.ID,
                    twitchhelix.ConditionStreamOnline{
                        BroadcasterUserID: "BROADCASTERS_ID",
                    },
                )
                if err != nil {
                    log.Fatal("SUBSCRIPTION INIT ERROR: ", err)
                }
            }
        case "notification":
            // Handle the event
            log.Println(event)
        }
}
```
