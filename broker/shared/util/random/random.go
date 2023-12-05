package random

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	alphabet string = "abcdefghijklmnopqrstuvwxyz"
	integer  string = "1234567890"
)

func RandomIP() string {
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
}

func RandomStrInt(n int) string {
	return fmt.Sprintf("%d", rand.Intn(n))
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
