package aquiladb

import (
	"math/rand"
)

func Test() string {
	return "Hello from module!!!"
}

func RundString() string {
	return RandStringBytes(10)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
