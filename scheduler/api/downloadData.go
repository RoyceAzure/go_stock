package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type URL_STOCK_DAY_AVG_ALL_DTO struct {
	Code            string `json:"Code"`
	StockName       string `json:"Name"`
	ClosingPrice    string `json:"ClosingPrice"`
	MonthlyAVGPRice string `json:"MonthlyAveragePrice"`
}

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
			fmt.Println(err)
			return nil, err
		}
		reqBody = bytes.NewBuffer([]byte(json_data))
	} else {
		reqBody = nil
	}
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		// 处理错误
		fmt.Println(err)
		return nil, err
	}

	if reqBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := client.Do(req)
	if err != nil {
		// 处理错误
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// 处理错误
		fmt.Println(err)
		return nil, err
	}

	return body, nil
}
