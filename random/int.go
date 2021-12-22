package random

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Int(min, max int) int {
	if min < 0 || min > max {
		panic("min must >= 0 && min <= max")
	}
	return rand.Intn(max-min+1) + min
}

func Int64(min, max int64) int64 {
	if min < 0 || min > max {
		panic("min must >= 0 && min <= max")
	}
	return rand.Int63n(max-min+1) + min
}
