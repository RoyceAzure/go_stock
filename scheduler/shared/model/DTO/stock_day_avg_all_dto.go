package dto

type StockDayAvgAllDTO struct {
	Code            string `json:"Code"`
	StockName       string `json:"Name"`
	ClosePrice      string `json:"ClosingPrice"`
	MonthlyAvgPrice string `json:"MonthlyAveragePrice"`
}
