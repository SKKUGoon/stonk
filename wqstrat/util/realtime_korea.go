package util

import (
	"fmt"
	"strings"
)

type DayTradeOrNightTrade string

const (
	// (To Korean) day trading
	USDayTrade DayTradeOrNightTrade = "D"

	// (To Korean) night trading - US real time
	USNightTrade DayTradeOrNightTrade = "R"
)

type OverseaExchangeRT string

const (
	NewYorkRT  OverseaExchangeRT = "NYS"
	NasdaqRT   OverseaExchangeRT = "NAS"
	AmexRT     OverseaExchangeRT = "AMS"
	TokyoRT    OverseaExchangeRT = "TSE"
	HongKongRT OverseaExchangeRT = "HKS"
	ShanghaiRT OverseaExchangeRT = "SHS"
	ShenZhenRT OverseaExchangeRT = "SZS"
	HochiminRT OverseaExchangeRT = "HSX"
	HanoiRT    OverseaExchangeRT = "HNX"
)

type rtExecRequestMessage struct {
	Header rtExecHeader `json:"header"`
	Body   rtExecBody   `json:"body"`
}

type rtExecHeader struct {
	ApprovalKey     string `json:"approval_key"`
	CustomerType    string `json:"custtype"`
	TransactionType string `json:"tr_type"`
	ContentType     string `json:"content-type"`
}

type rtExecBody struct {
	Input rtExecInput `json:"input"`
}

type rtExecInput struct {
	TransactionID string `json:"tr_id"`
	StockCode     string `json:"tr_key"`
}

func (c *KISClient) createRTExecHeader(register bool) (rtExecHeader, error) {
	var result rtExecHeader
	var tr string

	apprKey, err := c.OAuthWebsocket()
	if err != nil {
		return result, err
	}

	if register {
		tr = "1"
	} else {
		tr = "2"
	}

	result = rtExecHeader{
		ApprovalKey:     apprKey.ApprovalKey,
		CustomerType:    "P",
		TransactionType: tr,
		ContentType:     "utf-8",
	}

	return result, nil
}

func overseaSubscribeCode(dayOrNight DayTradeOrNightTrade, exchange OverseaExchangeRT, stockCode string) string {
	return strings.ToUpper(fmt.Sprintf("%s%s%s", dayOrNight, exchange, stockCode))
}

func (c *KISClient) SubscribeOversea(service, dayOrNight, exchange, stock string) error {
	// Check if KISClient exists
	client, ok := c.Streams[service]
	if !ok {
		return fmt.Errorf("no stream found for service %s", service)
	}

	// Create message
	header, err := c.createRTExecHeader(true)
	if err != nil {
		return err
	}

	msg := rtExecRequestMessage{
		Header: header,
		Body: rtExecBody{
			Input: rtExecInput{
				TransactionID: OverseaOrderExecutedTxID,
				StockCode:     overseaSubscribeCode(DayTradeOrNightTrade(dayOrNight), OverseaExchangeRT(exchange), stock),
			},
		},
	}

	// Send message
	if err = client.Conn.WriteJSON(msg); err != nil {
		return err
	}

	return nil
}

func (c *KISClient) Subscribe(service, stockCode string) error {
	// Check if KISClient exists
	client, ok := c.Streams[service]
	if !ok {
		return fmt.Errorf("no stream found for service %s", service)
	}

	// Create message
	header, err := c.createRTExecHeader(true) // Subscribe if true
	if err != nil {
		return err
	}

	msg := rtExecRequestMessage{
		Header: header,
		Body: rtExecBody{
			Input: rtExecInput{
				TransactionID: KorOrderExecutedTxID,
				StockCode:     stockCode,
			},
		},
	}

	// Send message
	if err = client.Conn.WriteJSON(msg); err != nil {
		return err
	}

	return nil
}

func (c *KISClient) Unsubscribe(service, stockCode string) error {
	// Check if KISClient exists
	client, ok := c.Streams[service]
	if !ok {
		return fmt.Errorf("no stream found for service %s", service)
	}

	// Create unsubscribing message
	header, err := c.createRTExecHeader(false) // Unsubscribe if false
	if err != nil {
		return err
	}

	msg := rtExecRequestMessage{
		Header: header,
		Body: rtExecBody{
			Input: rtExecInput{
				TransactionID: KorOrderExecutedTxID,
				StockCode:     stockCode,
			},
		},
	}

	if err = client.Conn.WriteJSON(msg); err != nil {
		return err
	}

	return nil
}
