package main

import (
	"crypto/pbkdf2"
	"crypto/sha256"
	"encoding/hex"
	"log/slog"
	"os"
	"phishing_backend/internal/infrastructure/presentation"
	"strconv"
	//_ "phishing_backend/internal/infrastructure/persistance" // include package so that init function is called
)

func main() {
	//hashPassword()
	presentation.SetupHttpServer()
}

func hashPassword() {
	// test=4c26730b8d9e68fe64bf1d029d36f5842def8b79ace77f8c3e0d61daf700ce00
	password := "test"
	hashedPassword, _ := pbkdf2.Key(sha256.New, password, []byte("z3aTemwpdxCpQswzhxKNNo2XDYPnTSUeDfeyyyodYD5phEr5TumiWpLTgyjCRmSN"), 1, 32)
	slog.Info(strconv.Itoa(len(hashedPassword)))
	slog.Info(hex.EncodeToString(hashedPassword))
}

func init() {
	setupDefaultLogger()
}

func setupDefaultLogger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}
