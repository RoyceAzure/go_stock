package api

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/RoyceAzure/go-stockinfo-schduler/util/constants"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

func Test_downloadDataAVG(t *testing.T) {
	var result []URL_STOCK_DAY_AVG_ALL_DTO
	res_body, err := SendRequest("GET", constants.URL_STOCK_DAY_AVG_ALL, nil)
	require.NoError(t, err)
	require.NotEmpty(t, res_body)

	err = json.Unmarshal(res_body, &result)
	require.NoError(t, err)
	require.NotEmpty(t, result)
}

func TestCreateSDAVGALL(t *testing.T) {
	var dtos []URL_STOCK_DAY_AVG_ALL_DTO
	byteData, err := SendRequest(constants.METHOD_GET,
		constants.URL_STOCK_DAY_AVG_ALL,
		nil)
	require.NoError(t, err)
	require.NotEmpty(t, byteData)

	json.Unmarshal(byteData, &dtos)
	require.NoError(t, err)
	context := context.Background()
	dto_len := len(dtos)
	cur_len := 0
	for _, dto := range dtos {
		entity, err := ConvertSDAVGALLDTO2E(&dto)
		if err != nil {
			log.Warn().
				Str("action", "ConvertSDAVGALLDTO2E").
				Any("parm", dto).
				Msg("Convert err, continue next")
			continue
		}
		res, err := testQueries.CreateSDAVGALL(context, *entity)
		require.NoError(t, err)
		require.NotEmpty(t, res)
		cur_len++
	}
	require.Equal(t, dto_len, cur_len)
}
