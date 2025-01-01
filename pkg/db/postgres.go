package db

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
    "os"
)

func Init() *sql.DB {
    dsn := os.Getenv("DATABASE_URL")
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        log.Fatal(err)
    }
    return db
}
