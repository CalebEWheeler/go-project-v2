package main

import (
	"github.com/CalebEWheeler/go-project-v2/database"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database.InitDatabase("rest_api")
}
