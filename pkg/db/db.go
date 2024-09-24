package db

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

var DB *sql.DB

func InitDB(connectionString string) {
    var err error
    DB, err = sql.Open("postgres", connectionString)
    if err != nil {
        log.Fatal("Database connection error:", err)
    }
}