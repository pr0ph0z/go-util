package util

import (
    "math/rand"
    "time"
)

func Rand(min int, max int) int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(max-min) + min
}