package constants

import "errors"

const (
	METHOD_POST           = "POST"
	METHOD_GET            = "GET"
	METHOD_PUT            = "PUT"
	METHOD_DELETE         = "DELETE"
	URL_STOCK_DAY_AVG_ALL = "https://openapi.twse.com.tw/v1/exchangeReport/STOCK_DAY_AVG_ALL"
	STR_ZERO              = "0.00"
)

var (
	ErrInValidatePreConditionOp = errors.New("invalid precondition of operation")
	ErrInternal                 = errors.New("internal error")
	ErrInvalidArgument          = errors.New("invaled argument")
)

var (
	ErrUserNotEsixts   = errors.New("user not exists or invalid password")
	ErrInvalidPassword = errors.New("wrong password")
)
