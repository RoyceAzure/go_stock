package util

import "sync"

func SliceBatchIterator[T any](ch chan<- []T, batchSize int, target []T, fiterFuncList []func([]T) []T) {
	i := 0
	if length := len(fiterFuncList); length != 0 {
		for i := 0; i < length; i++ {
			target = fiterFuncList[i](target)
		}
	}
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
