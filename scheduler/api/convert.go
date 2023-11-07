package api

import (
	repository "github.com/RoyceAzure/go-stockinfo-schduler/repository/sqlc"
	"github.com/RoyceAzure/go-stockinfo-schduler/util/constants"
)

func ConvertSDAVGALLDTO2E(dto *URL_STOCK_DAY_AVG_ALL_DTO) (*repository.CreateSDAVGALLParams, error) {
	if dto.ClosingPrice == "" {
		dto.ClosingPrice = constants.STR_ZERO
	}
	if dto.MonthlyAVGPRice == "" {
		dto.MonthlyAVGPRice = constants.STR_ZERO
	}
	return &repository.CreateSDAVGALLParams{
		Code:            dto.Code,
		StockName:       dto.StockName,
		ClosePrice:      dto.ClosingPrice,
		MonthlyAvgPrice: dto.MonthlyAVGPRice,
	}, nil
}
