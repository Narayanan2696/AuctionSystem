package service

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func URL() string {
	maxLen, _ := strconv.Atoi(os.Getenv("URL_LENGTH")) // assumed constant for generating url
	rand.Seed(time.Now().UnixNano())
	return String(rand.Intn(maxLen-1) + 1)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return "http://www.auction-arena.com/" + string(b)
}
