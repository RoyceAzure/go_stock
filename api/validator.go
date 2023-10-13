package api

import (
	"github.com/RoyceAzure/go-stockinfo-shared/utility"
	"github.com/go-playground/validator/v10"
)

var validSSO validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if t, ok := fieldLevel.Field().Interface().(string); ok {
		return utility.IsSupportedSSO(t)
	}
	return false
}
var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if t, ok := fieldLevel.Field().Interface().(string); ok {
		return utility.IsSupportedCurrencyType(t)
	}
	return false
}
