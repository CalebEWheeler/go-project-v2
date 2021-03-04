Thank you for visiting my REST API application. This application demonstrates my ability to translate my knowledge of Java to Go and create a basic REST API.

After cloning the repository to your local machine go through the SETUP steps before running the application.

SETUP: 

1. In the root of the project, create a new directory named 'config' and inside of that directory create a new file named 'config.go' 

2. Paste the lines 16-20 into 'config.go' and replace the values 'username' and 'password' in the declared const ()
        
  package config

  import "fmt"

  const (
	  username = "YOUR_MYSQL_USERNAME"
	  password = "YOUR_MYSQL_PASSWORD"
	  hostname = "127.0.0.1:3306"
  )

  func DSNString(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", username, password, hostname, dbName)
}

4. If you would like to change the name of the database and table to be created from "rest_api" and "person" you have the ability to do so. In database/database.go, navigate to the declared variables "dbName" and "tblName" and change the string values to the desired names. 
  * Also make sure to navigate to controllers/handlers.go and locate the declared variable tblName and make it's string value to be the same as tblName in database/database.go. 

  Here is an example if you would like the database name to be "office_personel" and the tblName to be "employee":  
  
    var dbName = "office_personel"
    var tblName = "employee"

  //If you would like to check, log into MySQL from the terminal and run these two Queries
    1. To see your newly created database - SHOW DATABASES; 
    2. To see your newly created 'employee' table - DESCRIBE employee;

  //If you would like to delete the database, because you would like to rename it more semantically. Log into MySQL from the terminal and run the Query: DROP DATABASE IF EXISTS office_personel;  

  
5. Next navigate to database/database.go and visit lines ( 61-63 ) to make changes to the set database settings I implemented. Your local machine may have other connections running, so make the changes accordingly. You can check how many connections your MySQL server can handle by running this MySQL Query: SHOW VARIABLES LIKE 'max_connections';


//Notes for Caleb getting into testing purposes...
1. make sure you pass in a string value for all key-value pairs in the body to test creating a person

ex: body {
  "name": "chester",
  "age": "33"
}
  


