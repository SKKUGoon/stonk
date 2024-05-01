package deribit

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	OptionREST = "https://www.deribit.com"
)

func structToMap(s interface{}) (map[string]string, error) {
	if s == nil {
		return nil, nil
	}

	// Change the json marshable struct into map[string]string
	var mapFromStruct map[string]string

	bstr, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bstr, &mapFromStruct)
	return mapFromStruct, err
}

// Performs GET http request without any query parameters
func get[T any](url string, query interface{}) (T, error) {
	var responseBody T

	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", OptionREST, url), nil)
	if err != nil {
		return responseBody, err
	}

	if qm, err := structToMap(query); err == nil {
		q := req.URL.Query()
		for k, v := range qm {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	} else {
		return responseBody, nil
	}

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return responseBody, err
	}
	defer response.Body.Close()

	if bytes, err := io.ReadAll(response.Body); err != nil {
		return responseBody, err
	} else {
		err = json.Unmarshal(bytes, &responseBody)
		return responseBody, err
	}
}
