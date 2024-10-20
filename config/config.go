// config/config.go
package config

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql" // MySQL driver
    "log"
)

var DB *sql.DB

// InitDB initializes the database connection
func InitDB() {
    var err error
    dsn := "root:Jai@1997@tcp(127.0.0.1:3306)/jai" // Update with your credentials

    // Open the database connection
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("Error opening database connection: %v", err)
    }

    // Test the connection
    err = DB.Ping()
    if err != nil {
        log.Fatalf("Error pinging database: %v", err)
    }

    fmt.Println("Database connection successfully established")
}
