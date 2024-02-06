package util

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/fatih/color"
	"github.com/gorilla/websocket"
)

const WebsocketUrl string = "/oauth2/Approval"

var websocketPath map[string]string = map[string]string{
	"korExec": ExecutedUrl,
}

var websocketMsgHandler map[string]func(string) error = map[string]func(string) error{
	"korExec": korExecMessage,
}

type OAuthWebsocketRequest struct {
	Auth
	GrantType string `json:"grant_type"`
}

type OAuthWebsocketResponse struct {
	ApprovalKey string `json:"approval_key"`
}

func (c *KISClient) OAuthWebsocket() (OAuthWebsocketResponse, error) {
	result := OAuthWebsocketResponse{}

	// Create request information
	body := OAuthWebsocketRequest{
		GrantType: "client_credentials",
		Auth:      c.UserInfo,
	}
	bstr, err := json.Marshal(body)
	if err != nil {
		return result, err
	}

	request, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s%s", OAuthForWebsocket, WebsocketUrl),
		bytes.NewReader(bstr),
	)
	if err != nil {
		return result, err
	}
	request.Header.Set("content-type", "application/json; utf-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("failed to get websocket access key: %v", err)
	}
	defer response.Body.Close()

	bytes, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(bytes, &result)
	return result, err
}

func (c *KISClient) StartStream(service string) error {
	// Generate websocket connection URL - for stream
	path, ok := websocketPath[service]
	if !ok {
		return errors.New(fmt.Sprintf("failed to get path for service %s", service))
	}

	// Create URL and create connection client
	u := url.URL{
		Scheme: RealTimeExecutedKorScheme,
		Host:   RealTimeExecutedKorHost,
		Path:   path,
	}
	color.Green(fmt.Sprintf("start service %s, connecting to %s", service, u.String()))

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}
	color.Green("websocket connection successful")

	// Create child context from the main context.
	// To prevent cancel func context leak, put `WithCancel` at the bottom
	ctx, cancel := context.WithCancel(c)

	c.Streams[service] = &KISStream{
		Context: ctx,
		Cancel:  cancel,
		Conn:    conn,
	}
	return nil
}

// Should be goroutine
func (c *KISClient) ReadFromSocket(service string) {
	client, ok := c.Streams[service]
	if !ok {
		log.Fatalf("failed to get stream for service %s", service)
	}

	for {
		select {
		case <-client.Context.Done():
			log.Printf("stopping stream for service %s\n", service)
			return
		default:
			_, message, err := client.Conn.ReadMessage()
			if err != nil {
				log.Printf("failed reading from websocket: %v", err)
			}
			websocketMsgHandler[service](string(message))
		}
	}
}

func (c *KISClient) CloseStream(service string) {
	if client, ok := c.Streams[service]; !ok {
		log.Fatalf("failed to get stream for service %s", service)
	} else {
		// Cancel context - stop read socket for loop
		client.Cancel()
		err := client.Conn.Close()
		if err != nil {
			log.Printf("graceful closing of websocket failed: %v", err)
		}
		color.Green(fmt.Sprintf("websocket connection for service %s closed", service))
	}
}
