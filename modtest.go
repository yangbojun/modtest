package modtest

import (
	"time"
	"math/rand"
)

func GetRand() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) + 1
}
