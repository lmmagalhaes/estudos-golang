package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)


const (
	dbDriver = "postgres"
	dbSource = "postgresql://go:go@localhost:5432/go?sslmode=disable"
)


func ConectaComBanco() *sql.DB {
    conn, err := sql.Open(dbDriver, dbSource)
    if err != nil {
      panic(err)
    }
    return conn
  }