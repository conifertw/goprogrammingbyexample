// To install Go-MySQL-Driver
// $ go get github.com/go-sql-driver/mysql
// It's installed in $GOPATH/pkg/darwin_amd64/github.com/go-sql-driver/mysql.q

package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    testConnection()
}

func testConnection() {
    // change database user and password
    db, err := sql.Open("mysql", string("user:password@tcp(127.0.0.1:3306)"))
    if err != nil {
        panic(err)
    }
    err = db.Ping() // test connection
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("connected")
    defer db.Close()
}

