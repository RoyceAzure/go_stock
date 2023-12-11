package util

import (
	"fmt"
	"net/mail"
	"regexp"

	"github.com/RoyceAzure/go-stockinfo/shared/util/constants"
	"github.com/shopspring/decimal"
)

var (
	isValidUsername = regexp.MustCompile(`^\w+$`).MatchString
	isValidFullname = regexp.MustCompile(`^[a-zA-Z\s]$`).MatchString
)

func IsSupportedSSO(sso string) bool {
	switch sso {
	case string(constants.MS), string(constants.GOOGLE), string(constants.FB), string(constants.AWS):
		return true
	}
	return false
}

func IsSupportedTransationType(trans_type string) bool {
	switch trans_type {
	case string(constants.BUY), string(constants.SELL):
		return true
	}
	return false
}

func IsSupportedCurrencyType(c_type string) bool {
	switch c_type {
	case string(constants.USD), string(constants.TW), string(constants.EU):
		return true
	}
	return false
}

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("string len must between %d-%d", minLength, maxLength)
	}
	return nil
}

// must not be zero value
func ValidateMustNotZeroInt64(value int64) error {
	if value == 0 {
		return fmt.Errorf("must not be zero value")
	}
	return nil
}

func ValidateStringToint64(value string) (int64, error) {
	if value == "" {
		return 0, fmt.Errorf("%w string is empty", constants.ErrInvalidArgument)
	}
	res, err := StringToInt64(value)
	return res, err
}

func ValidateStringToDecimal(value string) (decimal.Decimal, error) {
	if value == "" {
		return decimal.Zero, fmt.Errorf("%w value is empty", constants.ErrInvalidArgument)
	}
	res, err := decimal.NewFromString(value)
	return res, err
}

func ValidateMustGreateThenZero(value int64) error {
	if value <= 0 {
		return fmt.Errorf("must not be zero value")
	}
	return nil
}

func ValidateMustNotZeroInt(value int32) error {
	if value == 0 {
		return fmt.Errorf("must not be zero value")
	}
	return nil
}

// 3 <= string len <= 100 , only letters, digits, or underscore
func ValidateUsername(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if ok := isValidUsername(value); !ok {
		return fmt.Errorf("must contain only letters, digits, or underscore")
	}
	return nil
}

// 3 <= string len <= 100 , contain only letters or spaces
func ValidateFullname(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if ok := isValidFullname(value); !ok {
		return fmt.Errorf("must contain only letters or spaces")
	}
	return nil
}

// 3 <= string len <= 100
func ValidPassword(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	return nil
}

// 3 <= string len <= 100, only accept email format
func ValidEmail(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("email is invalid")
	}
	return nil
}

func ValidSSO(value string) error {
	if !IsSupportedSSO(value) {
		return fmt.Errorf("SSO is not supported")
	}
	return nil
}

func ValidEmailID(value int64) error {
	if value <= 0 {
		return fmt.Errorf("value must be postive")
	}
	return nil
}

func ValidSecretCode(value string) error {
	return ValidateString(value, 32, 128)
}
