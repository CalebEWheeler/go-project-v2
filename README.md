Thank you for visiting my REST API application. This application demonstrates my ability to translate my knowledge of Java to Go and create a basic REST API.

After cloning the repository to your local machine go through the SETUP steps before running the application.

SETUP: 

1. Start by logging into MySQL in the terminal and type this Query to create a database for the project. 'CREATE DATABASE IF NOT EXISTS database_name' replace "database_name" with whatever you would like to name your database for the application.

2. Next create a new directory named 'config' and inside of that directory create a new file named 'config.go' 

3. Paste the lines 16-20 into 'config.go' and replace these values in the string returned in MySQLConnectCred()
        
  package config

  import "fmt"

  const (
	  username = "YOUR_MYSQL_USERNAME"
	  password = "YOUR_MYSQL_PASSWORD"
	  hostname = "127.0.0.1:3306"
  )

  func DSNString(dbName string) string {
	  return fmt.Sprintf("%s:%s@tcp(%s)/", username, password, hostname)
  }

4. If you would like to change the name of the database to be created from 'person' you have the ability to do so. Navigate to 'database.go' to line 43 where you will see the func CreateDatabase called, change the second argument 'person' to what you would like the database name to be. 

  Here is an example if you would like the database name to be 'office_personel': CreateDatabase(db, "office_personel") 

  //If you would like to check, log into MySQL from the terminal and run the Query 'SHOW DATABASES;' to see your newly created database

  //If you would like to delete the database because you would like to rename it more semantically, log into MySQL from the terminal and run the Query 'DROP DATABASE IF EXISTS office_personel;'  

  

  


