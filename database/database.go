package database

//	TABLE OF CONTENTS:
//  ------------------
//	Line (): func to create table
//	Line (): func to create database
//	Line (): func to initialize the database and run all other functions for database setup

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/CalebEWheeler/go-project-v2/config"
	_ "github.com/go-sql-driver/mysql"
)

var dbName = "rest_api"
var tblName = "person"
var DB *sql.DB

func CreateDatabase(dbConn *sql.DB, dbName string) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := dbConn.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbName)
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

	//lines ( 39-47 ) validate if our DSN is correct
	dbConn, err := sql.Open("mysql", config.DSNString(""))

	if err != nil {
		log.Printf("Error %s when opening DB", err)
		return
	}

	defer dbConn.Close()
	fmt.Println("Successfully Connected to MySQL database")

	//line ( 50 ) will take the 'database connection' and a string for the desired 'database name' as arguments to create a new database.
	CreateDatabase(dbConn, dbName)

	//lines ( 53-58 ) will now validate if our DSN is correct in MySQL with our newly created database name passed as a string argument to config.DSNString()
	dbConn, err = sql.Open("mysql", config.DSNString(dbName))
	if err != nil {
		log.Printf("Error %s when opening DB", err)
		return
	}
	defer dbConn.Close()

	//This MySQL servers max connections is 151, I am implementing the following database settings on lines ( 61-63 ) in order to ensure there isn't an overload on the server.
	dbConn.SetMaxOpenConns(20)
	dbConn.SetMaxIdleConns(20)
	dbConn.SetConnMaxLifetime(time.Minute * 5)

	// lines ( 66-73 ) will utilize PingContext() to verify the actual connection to the database by pinging the DB.
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = dbConn.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return
	}
	log.Printf("Connected to DB %s successfully", dbName)

	// line ( 76 ) will take in the dbConn as the first argument and also a string value for the tblName as the second argument to create a new table with standard 'user' values.
	CreatePersonTable(dbConn, tblName)
}

func Connect() *sql.DB {
	db, err := sql.Open("mysql", config.DSNString(dbName))
	if err != nil {
		panic(err.Error())
	}

	return db
}
