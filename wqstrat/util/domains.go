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
	ExecutedUrl = "/tryitout/H0STCNT0"
)

const (
	ExecutedTxID = "H0STCNT0"
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

// Oversea
// REST "GET" Request with
// Header[T] and Body[U]
func overseaGETwHB[T any, U any, V any, X any](
	requestheader T,
	requestquery U,
	test bool,
	url string,
) (V, X, error) {
	var (
		resultHeader V
		resultBody   X

		headerMap map[string]string
		queryMap  map[string]string
	)

	req, err := http.NewRequest("GET", whereToRequest(test, url), nil)

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
		fmt.Println("generic", string(bytes))
		err = json.Unmarshal(bytes, &resultBody)
		return resultHeader, resultBody, err
	}
}
