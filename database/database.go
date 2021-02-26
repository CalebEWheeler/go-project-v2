package database

import (
	"database/sql"
	"fmt"

	"github.com/CalebEWheeler/go-project-v2/config"
	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase() {
	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", config.MySQLConnectCred())

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("Successfully Connected to MySQL database")

}
