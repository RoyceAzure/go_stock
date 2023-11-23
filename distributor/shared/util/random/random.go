package random

import (
	"fmt"
	"math/rand"
)

func RandomIP() string {
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
}

func RandomStrInt(n int) string {
	return fmt.Sprintf("%d", rand.Intn(n))
}
