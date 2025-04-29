//go:build integration

package integration_tests

import (
	"fmt"
	"net/http/httptest"
	"phishing_backend/internal/adapters/presentation"
	"testing"
)

var ts *httptest.Server

func TestMain(m *testing.M) {
	// Setup code
	fmt.Println("Setting up before tests")
	//loadEnvVariables()
	//startTestDB()
	//defer stopTestDB()
	ts = httptest.NewServer(presentation.NewServeMux())
	defer ts.Close()
	getDb().Exec("TRUNCATE TABLE users CASCADE")

	// Run the tests
	m.Run()

	// Teardown code
	fmt.Println("Cleaning up after tests")
}

//func loadEnvVariables() {
//	err := godotenv.Load("setup/integrationTests.env")
//	if err != nil {
//		fmt.Println("Error loading .env file", err)
//		os.Exit(1)
//	}
//}
//
//func startTestDB() error {
//	fmt.Println("Starting test DB...")
//	cmd := exec.Command("docker", "compose", "-f", "setup/docker-compose.integration.yml", "up", "-d", "--wait")
//	cmd.Stdout = os.Stdout
//	cmd.Stderr = os.Stderr
//	return cmd.Run()
//}
//
//func stopTestDB() error {
//	fmt.Println("Stopping test DB...")
//	cmd := exec.Command("docker", "compose", "-f", "setup/docker-compose.integration.yml", "down", "-v")
//	cmd.Stdout = os.Stdout
//	cmd.Stderr = os.Stderr
//	return cmd.Run()
//}
