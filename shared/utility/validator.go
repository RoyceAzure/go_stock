package utility

import (
	"fmt"
	"net/mail"
	"regexp"
)

type SSOType string
type CurrencyType string

const (
	MS     SSOType      = "MS"
	GOOGLE SSOType      = "GOOGLE"
	FB     SSOType      = "FB"
	AWS    SSOType      = "AWS"
	USD    CurrencyType = "USD"
	TW     CurrencyType = "TW"
	EU     CurrencyType = "EU"
)

var (
	isValidUsername = regexp.MustCompile(`^\w+$`).MatchString
	isValidFullname = regexp.MustCompile(`^[a-zA-Z\s]$`).MatchString
)

func IsSupportedSSO(sso string) bool {
	switch sso {
	case string(MS), string(GOOGLE), string(FB), string(AWS):
		return true
	}
	return false
}

func IsSupportedCurrencyType(c_type string) bool {
	switch c_type {
	case string(USD), string(TW), string(EU):
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
