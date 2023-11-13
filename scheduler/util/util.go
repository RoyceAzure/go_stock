package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

func SliceBatchIterator[T any](ch chan<- []T, batchSize int, target []T) {
	i := 0
	length := len(target)
	for i < length {
		end := i + batchSize
		if end > length {
			end = length
		}
		targetSlice := target[i:end]
		ch <- targetSlice
		i += batchSize
	}
	close(ch)
}

/*
將資料已batchSize的大小分配slice，並儲存到ch裡面
分配資料完畢時關閉ch
由於有關閉通道行為，所以只能由一個goroutine啟動
*/
func TaskDistributor[T any](ch chan<- []T, batchSize int, target []T, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(target); i += batchSize {
		end := i + batchSize
		if end > len(target) {
			end = len(target)
		}
		ch <- target[i:end]
	}
	close(ch)
}

/*
從unporcessed chan 接收[]資料
由processFunc 處理資料
最後儲存到porcessed chan

defer wg.Done()
*/
func TaskWorker[T any, T1 any](name string, unprocessed <-chan []T,
	porcessed chan<- T1,
	processFunc func(data T) (T1, error),
	errorFunc func(error),
	wg *sync.WaitGroup) {
	defer wg.Done()
	for dataBatch := range unprocessed {
		for _, data := range dataBatch {
			res, err := processFunc(data)
			if err != nil {
				if errorFunc != nil {
					errorFunc(err)
				}
				continue
			}
			porcessed <- res
		}
	}
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
