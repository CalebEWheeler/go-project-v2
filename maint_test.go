package main

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/CalebEWheeler/go-project-v2/config"
)

var dbConn *sql.DB

func setup(dbName string) {
	dbConn, _ := sql.Open("mysql", config.DSNString(dbName))
	dbConn.SetMaxIdleConns(5)

	fmt.Printf("\033[1;36m%s\033[0m", "> Setup completed\n")
}

func teardown() {
	_ = dbConn.Close()

	fmt.Printf("\033[1;36m%s\033[0m", "> Teardown completed\n")
	fmt.Printf("\n")
}

func TestMain(m *testing.M) {
	setup("rest_api")

	// EnsureTableExists()
	code := m.Run()
	// clearTable()
	teardown()
	os.Exit(code)
}
