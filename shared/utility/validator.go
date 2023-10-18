package utility

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
