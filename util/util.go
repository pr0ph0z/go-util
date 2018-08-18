package util

import (
	cRand "crypto/rand"
    mRand "math/rand"
    "time"
    "fmt"
)

func Rand(min int, max int) int {
    mRand.Seed(time.Now().UnixNano())
    return mRand.Intn(max-min) + min
}

func GenerateCode(n int) string {
    b := make([]byte, n)
    if _, err := cRand.Read(b); err != nil {
        panic(err)
    }
    s := fmt.Sprintf("%X", b)
    return s
}