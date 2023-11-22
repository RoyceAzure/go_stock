package jredis

import (
// "context"
// "testing"

// "github.com/RoyceAzure/go-stockinfo-schduler/service"
// "github.com/RoyceAzure/go-stockinfo-schduler/util/config"
// "github.com/stretchr/testify/require"

)

// func TestBulkInsertSPR(t *testing.T) {
// 	config, err := config.LoadConfig("../")
// 	require.NoError(t, err)
// 	r := NewJredis(config)
// 	require.NotEmpty(t, r)

// 	fakedataService := service.NewFakeSPRDataService(testDao)
// 	prorotype, err := fakedataService.GeneratePrototype()
// 	require.NoError(t, err)
// 	fakedataService.SetPrototype(prorotype)

// 	fakeDatas, err := fakedataService.GenerateFakeData()
// 	require.NoError(t, err)
// 	require.NotEmpty(t, r)
// 	require.Greater(t, len(fakeDatas), 0)

// 	key := "stock_real_time"

// 	err = r.BulkInsert(context.Background(), key, fakeDatas)
// 	require.NoError(t, err)
// }

// func TestGetSPR(t *testing.T) {
// 	config, err := config.LoadConfig("../")
// 	require.NoError(t, err)
// 	r := NewJredis(config)
// 	require.NotEmpty(t, r)

// 	fakedataService := service.NewFakeSPRDataService(testDao)
// 	prorotype, err := fakedataService.GeneratePrototype()
// 	require.NoError(t, err)
// 	fakedataService.SetPrototype(prorotype)

// 	fakeDatas, err := fakedataService.GenerateFakeData()
// 	require.NoError(t, err)
// 	require.NotEmpty(t, fakeDatas)
// 	require.Greater(t, len(fakeDatas), 0)

// 	key := "stock_real_time"
// 	ctx := context.Background()
// 	err = r.BulkInsert(ctx, key, fakeDatas)
// 	require.NoError(t, err)

// 	sprs, err := r.FindByID(ctx, key)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, sprs)
// 	require.Equal(t, len(sprs), len(fakeDatas))
// }
