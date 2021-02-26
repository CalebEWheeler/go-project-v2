Thank you for visiting my REST API application. This application demonstrates my ability to translate my knowledge of Java to Go and create a basic REST API.

After cloning the repository to your local machine go through the SETUP steps before running the application.

SETUP: 

1. Start by logging into MySQL in the terminal and type this Query to create a database for the project. 'CREATE DATABASE IF NOT EXISTS database_name' replace "database_name" with whatever you would like to name your database for the application.

2. Next create a new directory named 'config' and inside of that directory create a new file named 'config.go' 

3. Paste the lines 16-20 into 'config.go' and replace these values in the string returned in MySQLConnectCred()
        - ( 'root' with the name of your MySQL user ) 
        - ( 'password' with the password of your MySQL user )
        - ( 'database_name' with the database_name you created )
  
  package config

  func MySQLConnectCred() string {
	  return "root:password@tcp(127.0.0.1:3306)/database_name"
  }

4.   


