package repository

import (
	"database/sql"
	"fmt"
    _ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB
var dbUrl = "postgres://admin@localhost/silent_car_auction?sslmode=disable"
var dbDriver = "pgx"

func ensureDatabase() error {
    for db == nil {
        fmt.Println("Connecting to database...")
        var err error
        db, err = sql.Open(dbDriver, dbUrl)
        if err != nil {
            return err
        }
    }
    return nil
}