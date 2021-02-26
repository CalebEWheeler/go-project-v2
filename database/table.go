package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func CreatePersonTable(dbConn *sql.DB, tblName string) {
	query := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s(
		%s_id int primary key auto_increment,
		%s_name text, 
		%s_age int,
		created_at datetime default CURRENT_TIMESTAMP,
		updated_at datetime default CURRENT_TIMESTAMP
		)`, tblName, tblName, tblName, tblName)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	res, err := dbConn.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating %s table", err, tblName)
		return
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		return
	}
	log.Printf("Rows affected when creating table: %d", rows)
	return
}
