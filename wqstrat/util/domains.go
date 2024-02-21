package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	KoreaInvestREST     = "https://openapi.koreainvestment.com:9443"
	TestKoreaInvestREST = "https://openapivts.koreainvestment.com:29443"
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
	KorOrderExecutedUrl     = "/tryitout/H0STCNT0"
	OverseaOrderExecutedUrl = "/tryitout/HDFSCNT0"
)

const (
	// Body `tr_id` value for stream request

	KorOrderExecutedTxID     = "H0STCNT0"
	OverseaOrderExecutedTxID = "HDFSCNT0"
)

func whereToRequest(test bool, url string) string {
	switch test {
	case false:
		return fmt.Sprintf("%s%s", KoreaInvestREST, url)
	default:
		return fmt.Sprintf("%s%s", TestKoreaInvestREST, url)
	}
}

func structToMap(s interface{}) (map[string]string, error) {
	var mapFromStruct map[string]string

	bstr, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bstr, &mapFromStruct)
	return mapFromStruct, err
}

type OverseaGetRequestHeader struct {
	RESTAuth
	ContentType           string `json:"content-type,omitempty"`
	Authorization         string `json:"authorization"`
	PersonalSecurityKey   string `json:"personalseckey,omitempty"` // Essential for corporate client
	TransactionID         string `json:"tr_id"`
	TransactionContinued  string `json:"tr_cont,omitempty"`
	CustomerType          string `json:"custtype,omitempty"` // B for corporate client, P for individual
	SeqNo                 string `json:"seq_no,omitempty"`   // Essential for corporate client
	MacAddress            string `json:"mac_address,omitempty"`
	PhoneNumber           string `json:"phone_number,omitempty"`
	IPAddress             string `json:"ip_addr,omitempty"` // Essential for corporate client
	HashKey               string `json:"hashkey,omitempty"`
	GlobalTransactionUUID string `json:"gt_uid,omitempty"`
}

type OverseaGetResponseHeader struct {
	ContentType             string `json:"content-type"`
	TransactionID           string `json:"tr_id"`
	TransactionIsContinuous string `json:"tr_cont"`
	GlobalTransactionUUID   string `json:"gt_uid"`
}

// Oversea
// REST "GET" Request with
// RequestBody[T] and ResponseBody[U]
func overseaGETwHB[T any, U any](
	requestheader OverseaGetRequestHeader,
	requestquery T,
	test bool,
	url string,
) (OverseaGetResponseHeader, U, error) {
	var (
		resultHeader OverseaGetResponseHeader
		resultBody   U

		headerMap map[string]string
		queryMap  map[string]string
	)

	req, err := http.NewRequest("GET", whereToRequest(test, url), nil)
	if err != nil {
		return resultHeader, resultBody, err
	}

	// Set Header
	headerMap, err = structToMap(requestheader)
	if err != nil {
		return resultHeader, resultBody, err
	} else {
		for k, v := range headerMap {
			req.Header.Set(k, v)
		}
	}

	// Set Query Body
	queryMap, err = structToMap(requestquery)
	if err != nil {
		return resultHeader, resultBody, err
	} else {
		q := req.URL.Query()
		for k, v := range queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return resultHeader, resultBody, err
	}
	defer response.Body.Close()

	if bytes, err := io.ReadAll(response.Body); err != nil {
		return resultHeader, resultBody, err
	} else {
		err = json.Unmarshal(bytes, &resultBody)
		return resultHeader, resultBody, err
	}
}
