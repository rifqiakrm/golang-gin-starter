package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

const (
	five = 5
)

// CallerHeader define struct fir api call header
type CallerHeader struct {
	Key   string
	Value string
}

// CallAPI is a function for api call
func CallAPI(httpMethod, url string, headers []CallerHeader, payload interface{}) (*http.Response, error) {
	var req *http.Request
	var err error
	if payload != nil {
		payloadBuf := new(bytes.Buffer)
		if err := json.NewEncoder(payloadBuf).Encode(payload); err != nil {
			return nil, err
		}
		req, err = http.NewRequest(httpMethod, url, payloadBuf)
	} else {
		req, err = http.NewRequest(httpMethod, url, nil)
	}
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	for _, header := range headers {
		req.Header.Set(header.Key, header.Value)
	}

	client := &http.Client{
		Timeout: time.Minute * five,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
