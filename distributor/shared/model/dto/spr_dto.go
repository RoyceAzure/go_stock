package dto

import "time"

type StockPriceRealTimeDTO struct {
	StockCode    string    `json:"stock_code,omitempty"`
	StockName    string    ` json:"stock_name,omitempty"`
	TradeVolume  string    ` json:"trade_volume,omitempty"`
	TradeValue   string    ` json:"trade_value,omitempty"`
	OpenPrice    string    ` json:"open_price,omitempty"`
	HighestPrice string    ` json:"highest_price,omitempty"`
	LowestPrice  string    ` json:"lowest_price,omitempty"`
	ClosePrice   string    ` json:"close_price,omitempty"`
	Change       string    ` json:"change,omitempty"`
	Transaction  string    ` json:"transaction,omitempty"`
	TransTime    time.Time ` json:"trans_time,omitempty"`
}
