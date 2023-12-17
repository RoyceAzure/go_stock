package constants

import "errors"

type SSOType string
type CurrencyType string

const (
	ForeignKeyViolation              = "foreign_key_violation"
	UniqueViolation                  = "unique_violation"
	MS                  SSOType      = "MS"
	GOOGLE              SSOType      = "GOOGLE"
	FB                  SSOType      = "FB"
	AWS                 SSOType      = "AWS"
	USD                 CurrencyType = "USD"
	TW                  CurrencyType = "TW"
	EU                  CurrencyType = "EU"
	DEFAULT_PAGE                     = 1
	DEFAULT_PAGE_SIZE                = 10
	SELL                             = "sell"
	BUY                              = "buy"
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
