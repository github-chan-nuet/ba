package services

import (
	"crypto/pbkdf2"
	"crypto/sha256"
	"log/slog"
	"os"
	"strconv"
)

var (
	pbkdf2Salt = os.Getenv("PHBA_PBKDF2_SALT")
	pbkdf2Iter = os.Getenv("PHBA_PBKDF2_ITER")
)

func HashPassword(password string) []byte {
	// https://de.wikipedia.org/wiki/PBKDF2
	hashedPassword, err := pbkdf2.Key(sha256.New, password, []byte(pbkdf2Salt), mustAtoi(pbkdf2Iter), 32)
	if err != nil {
		slog.Error("Could not hash password", "err", err)
		panic("Could not hash password:" + err.Error())
	}
	return hashedPassword
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic("could not convert string to int, string: " + s)
	}
	return i
}
