package main

/***

This program is example of
1) using powclient from public pkg
2) how to solve pow challenge itself (brute-force on random strings)

***/

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"

	"github.com/MedvedewEM/pow/pkg/powclient"
	"github.com/MedvedewEM/pow/pkg/slices"
)

const (
	MAX_RETRY_ON_SAME_LENGTH = 26
	MAX_RETRY_TO_STOP        = 10000

	errUnableToSolve = "errUnableToSolve"
)

func main() {
	client := powclient.New("http://pow_http_server:8080")

	id, suffix, err := client.Please()
	if err != nil {
		fmt.Println("No please:", err)
		return
	}

	answer, err := TryToSolve(suffix)
	if err != nil {
		fmt.Println("No solution or error:", err)
		return
	}

	word, err := client.WisdomWord(id, answer)
	if err != nil {
		fmt.Println("No wisdom word", err)
		return
	}

	fmt.Println(word)
}

// TryToSolve tries to find string which has appropriate hash's ending
func TryToSolve(suffix string) (string, error) {
	expected, err := hex.DecodeString(suffix)
	if err != nil {
		return "", errors.New(errUnableToSolve)
	}

	currentLen := 1
	for currentLen < MAX_RETRY_TO_STOP {
		currentRetry := 0
		for currentRetry < MAX_RETRY_ON_SAME_LENGTH {
			randomStr := randStringRunes(currentLen)

			actual := sha256LastBytes(randomStr, len(expected))

			if slices.Equal(expected, actual) {
				return randomStr, nil
			}

			currentRetry++
		}

		currentLen++
	}

	return "", errors.New(errUnableToSolve)
}

// sha256LastBytes returns last N bytes of sha256 hash of given string
func sha256LastBytes(str string, N int) []byte {
	hasher := sha256.New()
	hasher.Write([]byte(str))
	actualBytes := hasher.Sum(nil)

	return actualBytes[len(actualBytes)-N:]
}

// randStringRunes generates random string by given length
func randStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}
