/*
Readme:

* Enable TCP/IP Network protocol for SQL Server
  Navigate to Computer Management >  Services and Applications > SQL Server Configuration Manager
  > SQL Server Network Configuration
* Enable TCP/IP
* Restart SQL Server (MSSQLSERVER)

*/ 
package main;

import(
"fmt"
"gorm.io/gorm"
"gorm.io/driver/sqlserver"
)

// How to install gorm? go get -u gorm.io/gorm

type User struct{
	gorm.Model
	Email string    `json:"email" gorm:"unique"` 
   }


 func main() { 

// for windows authentication
connectionString := "sqlserver://:@127.0.0.1:1433?database=GoLangDB"
// for SQL Server authentication
//connectionString := "sqlserver://username:password@host:port?database=nameDB"

	db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
	   if err != nil {
		   fmt.Println("failed to connect database") 
		   panic("failed to connect database")
	   }
	
	   // Migrate the schema
	   db.Migrator().CreateTable(&User{})
}

