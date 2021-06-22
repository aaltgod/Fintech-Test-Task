package tools

import (
	"math/rand"
	"time"
)

func RandomID(IDlen int) string {

	rand.Seed(time.Now().UnixNano())

	alph := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	result := make([]rune, IDlen)
	for i := range result {
		result[i] = alph[rand.Intn(len(alph))]
	}

	return string(result)
}
