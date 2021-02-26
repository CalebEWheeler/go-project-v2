package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/CalebEWheeler/go-project-v2/config"
	_ "github.com/go-sql-driver/mysql"
)

func CreateDatabase(db *sql.DB, dbName string) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbName)
	if err != nil {
		log.Printf("Error %s when creating DB\n", err)
		return
	}
	no, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when fetching rows", err)
		return
	}
	log.Printf("rows affected %d\n", no)
}

func InitDatabase() {

	db, err := sql.Open("mysql", config.DSNString(""))

	if err != nil {
		log.Printf("Error %s when opening DB", err)
		return
	}

	defer db.Close()

	fmt.Println("Successfully Connected to MySQL database")

	CreateDatabase(db, "people")
}
