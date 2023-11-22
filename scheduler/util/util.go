package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/jackc/pgx/v5/pgtype"
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

func NumericToFloat64(n pgtype.Numeric) (float64, error) {
	if !n.Valid {
		return 0, errors.New("numeric value not present")
	}

	float8, err := n.Float64Value()
	if err != nil {
		return 0, err
	}
	return float8.Float64, nil
}

func float64ToNumeric(f float64) (pgtype.Numeric, error) {
	// 創建一個新的 Numeric 值
	n := pgtype.Numeric{}

	// 使用 Set 方法設置 float64 值
	if err := n.Scan(f); err != nil {
		// 如果轉換有錯誤，返回錯誤
		return pgtype.Numeric{}, err
	}

	return n, nil
}

func StringToNumeric(value string) (pgtype.Numeric, error) {
	var result pgtype.Numeric
	err := result.Scan(value)
	return result, err
}

func Float64ToString(value float64) string {
	// 第二個參數指定格式，'f' 表示不使用指數形式
	// 第三個參數指定小數點後的位數
	// 第四個參數用於指定要轉換的浮點型號，64 表示 float64
	return strconv.FormatFloat(value, 'f', 2, 64)
}

func ReadJsonFile(path string) ([]byte, error) {
	if path == "" {
		return nil, fmt.Errorf("read json, path is empty")
	}
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read json, get err : %w", err)
	}
	return file, nil
}

func WriteJsonFile(path string, data any) error {
	if path == "" {
		return fmt.Errorf("write json, path is empty")
	}
	var byteData []byte
	var err error
	var ok bool
	if byteData, ok = data.([]byte); !ok {
		byteData, err = json.Marshal(data)
		if err != nil {
			return fmt.Errorf("write json file, marshal json get err : %w", err)
		}
	}

	err = os.WriteFile(path, byteData, 0644)
	if err != nil {
		return fmt.Errorf("write json file, get err : %w", err)
	}

	return nil
}
