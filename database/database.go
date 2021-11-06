package database

import (
    "database/sql"
    "fmt"

    _ "github.com/lib/pq"
)

const DriverPostgres = "postgres"

type Connection struct {
    Host     string
    Port     int
    Driver   string
    Database string
    Username string
    Password string
    SSLMode  string
}

var db *sql.DB

func GetDatabase(connection Connection) (*sql.DB, error) {
    if db != nil {
        return db, nil
    }

    var err error
    db, err = sql.Open(connection.Driver, connection.getDatasourceName())
    if err != nil {
        return nil, err
    }
    if err = db.Ping(); err != nil {
        return nil, err
    }

    return db, nil
}

func (dc Connection) getDatasourceName() string {
    return fmt.Sprintf("%s:%s@%s:%d/%s?sslmode=%s", dc.Username, dc.Password, dc.Host, dc.Port, dc.Database, dc.SSLMode)
}
