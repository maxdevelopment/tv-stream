package utils

import (
	"math/rand"
	"crypto/md5"
	"time"
	"encoding/hex"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)

	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	hash := md5.New()
	text := time.Now().String()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil)) + string(b)
}
