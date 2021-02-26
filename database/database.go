package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/CalebEWheeler/go-project-v2/config"
	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase() {
	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", config.DSNString(""))

	if err != nil {
		log.Printf("Error %s when opening DB", err)
		return
	}

	defer db.Close()

	fmt.Println("Successfully Connected to MySQL database")

}
