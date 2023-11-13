package api

// func Test_downloadDataAVG(t *testing.T) {
// 	var result []URL_STOCK_DAY_AVG_ALL_DTO
// 	res_body, err := SendRequest("GET", constants.URL_STOCK_DAY_AVG_ALL, nil)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, res_body)

// 	err = json.Unmarshal(res_body, &result)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, result)
// }

// func TestCreateSDAVGALL(t *testing.T) {
// 	var dtos []URL_STOCK_DAY_AVG_ALL_DTO
// 	byteData, err := SendRequest(constants.METHOD_GET,
// 		constants.URL_STOCK_DAY_AVG_ALL,
// 		nil)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, byteData)

// 	json.Unmarshal(byteData, &dtos)
// 	require.NoError(t, err)
// 	ctx := context.Background()
// 	dto_len := len(dtos)
// 	cur_len := 0
// 	for _, dto := range dtos {
// 		entity, err := ConvertSDAVGALLDTO2E(dto)
// 		if err != nil {
// 			log.Warn().
// 				Str("action", "ConvertSDAVGALLDTO2E").
// 				Any("parm", dto).
// 				Msg("Convert err, continue next")
// 			continue
// 		}
// 		res, err := testDao.CreateSDAVGALL(ctx, entity)
// 		require.NoError(t, err)
// 		require.NotEmpty(t, res)
// 		cur_len++
// 	}
// 	require.Equal(t, dto_len, cur_len)
// }

// func TestBulkCreateSDAVGALL(t *testing.T) {
// 	var dtos []URL_STOCK_DAY_AVG_ALL_DTO
// 	var entities []repository.BulkInsertDAVGALLParams
// 	byteData, err := SendRequest(constants.METHOD_GET,
// 		constants.URL_STOCK_DAY_AVG_ALL,
// 		nil)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, byteData)

// 	json.Unmarshal(byteData, &dtos)
// 	require.NoError(t, err)
// 	ctx := context.Background()
// 	dto_len := len(dtos)
// 	totalStartTime := time.Now()
// 	ch := make(chan []URL_STOCK_DAY_AVG_ALL_DTO)
// 	batchSize := 2000
// 	var handler_len int64
// 	go util.SliceBatchIterator(ch, batchSize, dtos)
// 	for dtos := range ch {
// 		for _, dto := range dtos {
// 			entity, err := ConvertSDAVGALLDTO2BulkEntity(dto)
// 			if err != nil {
// 				log.Warn().
// 					Str("action", "ConvertSDAVGALLDTO2E").
// 					Any("parm", dto).
// 					Msg("Convert err, continue next")
// 				continue
// 			}
// 			entities = append(entities, entity)
// 		}
// 		batchInsertTime := time.Now()
// 		res, err := testDao.BulkInsertDAVGALL(ctx, entities)
// 		require.WithinDuration(t, time.Now(), batchInsertTime, time.Millisecond*2)
// 		require.NoError(t, err)
// 		require.NotEmpty(t, res)
// 		handler_len += res
// 	}
// 	require.EqualValues(t, handler_len, dto_len)
// 	require.WithinDuration(t, time.Now(), totalStartTime, time.Second)
// }

// func TestBulkCreateSDAVGALLV2(t *testing.T) {
// 	var dtos []URL_STOCK_DAY_AVG_ALL_DTO
// 	var entities []repository.BulkInsertDAVGALLParams
// 	byteData, err := SendRequest(constants.METHOD_GET,
// 		constants.URL_STOCK_DAY_AVG_ALL,
// 		nil)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, byteData)

// 	json.Unmarshal(byteData, &dtos)
// 	require.NoError(t, err)
// 	ctx := context.Background()
// 	dto_len := len(dtos)
// 	totalStartTime := time.Now()
// 	ch := make(chan []URL_STOCK_DAY_AVG_ALL_DTO)
// 	batchSize := 2000
// 	var handler_len int64
// 	go util.SliceBatchIterator(ch, batchSize, dtos)
// 	for dtos := range ch {
// 		for _, dto := range dtos {
// 			entity, err := ConvertSDAVGALLDTO2BulkEntity(dto)
// 			if err != nil {
// 				log.Warn().
// 					Str("action", "ConvertSDAVGALLDTO2E").
// 					Any("parm", dto).
// 					Msg("Convert err, continue next")
// 				continue
// 			}
// 			entities = append(entities, entity)
// 		}
// 		batchInsertTime := time.Now()
// 		res, err := testDao.BulkInsertDAVGALL(ctx, entities)
// 		require.WithinDuration(t, time.Now(), batchInsertTime, time.Millisecond*2)
// 		require.NoError(t, err)
// 		require.NotEmpty(t, res)
// 		handler_len += res
// 	}
// 	require.EqualValues(t, handler_len, dto_len)
// 	require.WithinDuration(t, time.Now(), totalStartTime, time.Second)
// }
