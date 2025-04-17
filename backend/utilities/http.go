package utilities

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func HTTPRequest(method string, route string, body []byte, headers map[string]string) (*http.Response, error) {
	url := fmt.Sprintf("http://localhost:%s/api%s", os.Getenv("PORT"), route)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
