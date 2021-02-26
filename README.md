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
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

4. If you would like to change the name of the database to be created from 'person' you have the ability to do so. Navigate to 'database.go' to line 49 where you will see the func CreateDatabase() called, change the second argument "person" to what you would like the database name to be. 

  Here is an example if you would like the database name to be "office_personel": CreateDatabase(db, "office_personel") 

  //If you would like to check, log into MySQL from the terminal and run the Query: SHOW DATABASES; to see your newly created database

  //If you would like to delete the database, because you would like to rename it more semantically. Log into MySQL from the terminal and run the Query: DROP DATABASE IF EXISTS office_personel;  

  * Also make sure to visit lines ( 61-63 ) to make changes to the set database settings I implemented. Your local machine may have other connections running, so make the changes accordingly. You can check how many connections your MySQL server can handle by running this MySQL Query: SHOW VARIABLES LIKE 'max_connections';

  

  


