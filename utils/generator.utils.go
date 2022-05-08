package utils

import (
	cryptoRand "crypto/rand"
	"fmt"
	"io"
	"math/rand"
	"time"
)

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func GenerateTrxID(prefix string) string {
	res := prefix
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(9999-1000) + 10000
	res = res + time.Now().Format("20060102") + "/" + fmt.Sprint(num)

	return res
}

func GenerateExternalID(prefix string) string {
	res := prefix + fmt.Sprint(time.Now().Unix())

	return res
}

func GenerateOTP(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(cryptoRand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
