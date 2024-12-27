package util

import (
	"strings"
	"time"

	"math/rand"

	"github.com/bxcodec/faker/v3"
)

const ALPHABET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var random *rand.Rand

func init() {
	seed := int64(time.Now().UnixNano())
	random = rand.New(rand.NewSource(seed))
}

func RandomInt(min, max int64) int64 {
	return min + random.Int63n(max-min+1)
}

func RandomString(n int) string {
	var builder strings.Builder
	k := len(ALPHABET)

	for i := 0; i < n; i++ {
		c := ALPHABET[random.Intn(k)]
		builder.WriteByte(c)
	}

	return builder.String()
}

func RandomPassword() string {
	n := int(RandomInt(8, 20))
	return RandomString(n)
}

func RandomOwner() string {
	return faker.Username()
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[random.Intn(n)]
}

func RandomEmail() string {
	return faker.Email()
}
