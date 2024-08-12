package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type RequestOptions struct {
	Method          string
	URL             string
	Headers         map[string]string
	QueryParams     map[string]string
	Body            []byte
	TimeoutInSecond time.Duration
}

// content type
const (
	CONTENT_TYPE_FORM_URLENCODED  = "application/x-www-form-urlencoded"
	CONTENT_TYPE_APPLICATION_JSON = "application/json"
)

var responseHeaderKey = []string{
	"Correlationid", //Brigate
}

func SendHttpRequest(ctx context.Context, request RequestOptions) (statusCode int, respBody []byte, err error) {
	var resp *http.Response
	var req *http.Request

	if request.TimeoutInSecond < 1 {
		request.TimeoutInSecond = 60
	}

	client := &http.Client{
		Timeout: request.TimeoutInSecond * time.Second,
	}

	if len(request.QueryParams) > 0 {
		values := url.Values{}
		for key, value := range request.QueryParams {
			values.Add(key, value)
		}
		request.URL += "?" + values.Encode()
	}

	defer func() {
		responseHeader := make(map[string]interface{})

		if resp != nil && resp.Header != nil {
			for i := 0; i < len(responseHeaderKey); i++ {
				for key, value := range resp.Header {
					if responseHeaderKey[i] == key {
						responseHeader[key] = value
						break
					}
				}
			}
		}
	}()

	req, err = http.NewRequest(request.Method, request.URL, bytes.NewBuffer(request.Body))
	if err != nil {
		return
	}

	if len(request.Headers) != 0 {
		for k, v := range request.Headers {
			req.Header.Set(k, v)
		}
	}

	resp, err = client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	statusCode = resp.StatusCode

	if statusCode >= 500 {
		err = errors.New("Server Error")
	} else if statusCode >= 400 {
		err = errors.New("Request Error")
	}

	return
}

func jsonStringToMap(data []byte) map[string]interface{} {
	var result map[string]interface{}

	err := json.Unmarshal(data, &result)
	if err != nil {
		return nil
	}
	return result
}
