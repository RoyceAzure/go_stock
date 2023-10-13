package utility

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet string = "abcdefghijklmnopqrstuvwxyz@."

var TransactionType = [2]string{"Buy", "Sell"}

var CurrencyTypeSlice = []CurrencyType{TW, USD, EU}

var SSOTypeSlice = []SSOType{MS, GOOGLE, AWS, FB}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomFloat(max float64) float64 {
	return rand.Float64() * max
}

func RandomTransactionType() string {
	return TransactionType[rand.Intn(len(TransactionType))]
}

func RandomSSOType() SSOType {
	return SSOTypeSlice[rand.Intn(len(SSOTypeSlice))]
}

func RandomCurrencyType() CurrencyType {
	return CurrencyTypeSlice[rand.Intn(len(CurrencyTypeSlice))]
}

func RandomSSOTypeStr() string {
	return string(RandomCurrencyType())
}

func RandomCurrencyTypeStr() string {
	return string(RandomCurrencyType())
}

func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}
