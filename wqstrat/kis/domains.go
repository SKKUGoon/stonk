package kis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	KoreaInvestREST     = "https://openapi.koreainvestment.com:9443"
	TestKoreaInvestREST = "https://openapivts.koreainvestment.com:29443"
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

type OverseaRequestHeader struct {
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

type OverseaResponseHeader struct {
	ContentType             string `json:"content-type"`
	TransactionID           string `json:"tr_id"`
	TransactionIsContinuous string `json:"tr_cont"`
	GlobalTransactionUUID   string `json:"gt_uid"`
}

type OverseaGetResponseBodyBase struct {
	ReturnCode  string `json:"rt_cd"` // 0 if success
	MessageCode string `json:"msg_cd"`
	Message     string `json:"msg1"`
}

// Oversea
// REST "GET" Request with
// RequestBody[T] and ResponseBody[U]
func overseaGETwHB[T any, U any](
	requestheader OverseaRequestHeader,
	requestquery T,
	test bool,
	url string,
) (OverseaResponseHeader, U, error) {
	var (
		resultHeader OverseaResponseHeader
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

func overseaPOSTwHB[T any, U any](
	requestheader OverseaRequestHeader,
	requestbody T,
	test bool,
	url string,
) (OverseaResponseHeader, U, error) {
	var (
		resultHeader OverseaResponseHeader
		resultBody   U

		headerMap map[string]string
	)

	bstr, err := json.Marshal(requestbody)
	if err != nil {
		return resultHeader, resultBody, err
	}

	req, err := http.NewRequest("POST", whereToRequest(test, url), bytes.NewReader(bstr))
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
