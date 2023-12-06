package util

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

// type DataMakerinterface interface{
// 	makeData[T any] (int, T) []T
// }

// type SliceDataMaker[T any] struct{}

// func (s SliceDataMaker[T]) makeData(dataLegnth int, t T) []T {
// 	return make([]T, dataLegnth)
// }

func makeData[T any](length int) []T {
	return make([]T, length)
}

func TestSliceIterator(t *testing.T) {
	var num []int
	length := 100
	var totalResult []int
	for i := 0; i < length; i++ {
		num = append(num, i)
	}
	batchSize := 10
	ch := make(chan []int)
	go SliceBatchIterator(ch, batchSize, num)

	resultCount := 0
	for result := range ch {
		resultCount++
		fmt.Printf("%v", result)
		totalResult = append(totalResult, result...)
	}
	fmt.Printf("resultCount %d", resultCount)
	require.Equal(t, length, len(totalResult))
}

/*
TODO 不要等到資料都處理完才寫入DB
*/
func TestSliceIterAndBroker(t *testing.T) {
	var num []int
	var wg sync.WaitGroup
	length := 100
	var resultList []float32
	for i := 0; i < length; i++ {
		num = append(num, i)
	}
	batchSize := 15
	unporcessed := make(chan []int)
	processed := make(chan float32)
	wg.Add(3)
	go TaskDistributor(unporcessed, batchSize, num, &wg)
	go TaskWorker("broker1", unporcessed, processed, processfun, nil, &wg)
	go TaskWorker("broker2", unporcessed, processed, processfun, nil, &wg)

	// go func() {
	// 	wg.Wait()
	// 	close(porcessed)
	// }()
	insertCount := 0
	go func() {
		wg.Wait()
		close(processed)
	}()
	for data := range processed {
		resultList = append(resultList, data)
		if len(resultList)%batchSize == 0 {
			insertCount += len(resultList)
			resultList = make([]float32, 0)
		}
	}
	if len(resultList) > 0 {
		insertCount += len(resultList)
	}
	require.Equal(t, length, insertCount)
}

func processfun(data int) (res float32, err error) {
	res = float32(data) + 0.01
	return res, nil
}

func TestWriteJsonFile(t *testing.T) {
	byteData, err := ReadJsonFile("./STOCK_DAY_ALL.json")
	require.NoError(t, err)
	err = WriteJsonFile("./STOCK_DAY_ALL.json", byteData)
	require.NoError(t, err)
}
