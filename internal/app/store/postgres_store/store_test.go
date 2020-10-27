package postgres_store_test

import (
	"fmt"
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {

	databaseURL = os.Getenv("PASSWORD")
	fmt.Println("->", os.Environ())
	if databaseURL == "" {
		databaseURL = "port=5432 user=postgres password=12345 dbname=buyer_exp sslmode=disable"
	}

	fmt.Println("->", databaseURL)
	os.Exit(m.Run())
}
