package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"phishing_backend/internal/adapters/presentation/api"
)

var (
	ErrorInternal = errors.New("internal error")
)

func getError(level int) error {
	level1Err := fmt.Errorf("[getData] level 1 error: %w", ErrorInternal)
	if level == 1 {
		return level1Err
	}
	if level == 2 {
		return fmt.Errorf("[getData] level 2 error: %w", level1Err)
	}

	return ErrorInternal
}

func main() {
	c()
}

func a() {
	err := getError(0)
	if errors.Is(err, ErrorInternal) {
		fmt.Printf("is error internal: %v\n", err)
	}
	fmt.Printf("unwrapped error: %v\n", err.Error())

	fmt.Printf("---\n")

	err = getError(1)
	if errors.Is(err, ErrorInternal) {
		fmt.Printf("is error internal: %v\n", err)
	}
	fmt.Printf("unwrapped error: %v\n", err.Error())

	fmt.Printf("---\n")

	err = getError(2)
	if errors.Is(err, ErrorInternal) {
		fmt.Printf("is error internal: %v\n", err)
	}
	fmt.Printf("unwrapped error: %v\n", err.Error())
}

func b() {
	errA := errors.New("error a")
	errB := errors.New("error b")
	errCombined := errors.Join(errA, errB)
	if errors.Is(errCombined, errA) {
		fmt.Println("is errA")
	}
	if errors.Is(errCombined, errB) {
		fmt.Println("is errB")
	}
	fmt.Println(errCombined.Error())
}

func c() {
	var user api.UserPostModel
	jsonData := `{ "a": "a", "b": 1 }`
	err := json.Unmarshal([]byte(jsonData), &user)
	fmt.Println(user)
	fmt.Println(err)
}
