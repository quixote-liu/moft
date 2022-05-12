package util

import (
	"math/rand"
	"strconv"
	"time"
)

func RandomInt64() string {
	rand.Seed(time.Now().Unix())
	return strconv.FormatInt(rand.Int63(), 10)
}

func RandomInt32() string {
	rand.Seed(time.Now().Unix())
	return strconv.FormatInt(int64(rand.Int31()), 10)
}
