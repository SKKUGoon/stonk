package util

import "encoding/json"

func structToMap(s interface{}) (map[string]string, error) {
	var mapFromStruct map[string]string

	bstr, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bstr, &mapFromStruct)
	return mapFromStruct, err
}
