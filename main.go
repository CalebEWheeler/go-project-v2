package main

import (
	"net/http"

	"github.com/CalebEWheeler/go-project-v2/controllers"
	_ "github.com/go-sql-driver/mysql"
)

var dbName = "rest_api"
var tblName = "person"

func main() {
	http.Handle("/", controllers.SetupRoutes())
}
