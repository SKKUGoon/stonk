package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/fatih/color"
	"github.com/gorilla/websocket"
)

// Websocket
const (
	RealTimeExecutedKor       = "ws://ops.koreainvestment.com:21000"
	RealTimeExecutedKorScheme = "ws"
	RealTimeExecutedKorHost   = "ops.koreainvestment.com:21000"

	TestRealTimeExecutedKor       = "ws://ops.koreainvestment.com:31000"
	TestRealTimeExecutedKorScheme = "ws"
	TestRealTimeExecutedKorHost   = "ops.koreainvestment.com:31000"
)

const (
	// Body `tr_id` value for stream request

	KorOrderExecutedTxID     = "H0STCNT0"
	OverseaOrderExecutedTxID = "HDFSCNT0"
)

const (
	WebsocketUrl = "/oauth2/Approval"

	KorOrderExecutedUrl     = "/tryitout/H0STCNT0"
	OverseaOrderExecutedUrl = "/tryitout/HDFSCNT0"
)

var websocketPath map[string]string = map[string]string{
	// Korean market service
	"korExec": KorOrderExecutedUrl,

	// Oversea market service
	"overseaExec": OverseaOrderExecutedUrl,
}

var websocketMsgHandler map[string]func(string) error = map[string]func(string) error{
	"korExec": korExecMessage,

	"overseaExec": overseaExecMessage,
}

type OAuthWebsocketRequest struct {
	WebsocketAuth
	GrantType string `json:"grant_type"`
}

type OAuthWebsocketResponse struct {
	ApprovalKey string `json:"approval_key"`
}

func (c *KISClient) OAuthWebsocket() (OAuthWebsocketResponse, error) {
	result := OAuthWebsocketResponse{}

	// Create request information
	body := OAuthWebsocketRequest{
		GrantType:     "client_credentials",
		WebsocketAuth: c.UserInfoWS,
	}
	bstr, err := json.Marshal(body)
	if err != nil {
		return result, err
	}

	request, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s%s", KoreaInvestREST, WebsocketUrl),
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
		return fmt.Errorf("failed to get path for service %s", service)
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

			// Attach handlers to the string to process messages
			if fn, ok := websocketMsgHandler[service]; ok {
				fn(string(message))
			} else {
				fmt.Printf("no handlers for service %s", service)
				fmt.Println(string(message))
			}
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
