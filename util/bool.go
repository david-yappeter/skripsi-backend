package util

import (
	"math/rand"
	"time"
)

func RandBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}

func BoolP(v bool) *bool {
	return &v
}
