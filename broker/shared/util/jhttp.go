package util

import (
	"bytes"
	"encoding/json"
	"io"
	"net"
	"net/http"
)

func SendRequest(method string, url string, data any) ([]byte, error) {
	client := &http.Client{}
	var reqBody io.Reader

	if method == "POST" || method == "PUT" {
		reqBody = &bytes.Buffer{}
	}

	if data != nil {
		json_data, err := json.Marshal(data)
		if err != nil {
			// 处理错误
			return nil, err
		}
		reqBody = bytes.NewBuffer([]byte(json_data))
	} else {
		reqBody = nil
	}
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, err
	}

	if reqBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func IsValidIP(ip string) bool {
	parsedIP := net.ParseIP(ip)
	return parsedIP != nil
}
