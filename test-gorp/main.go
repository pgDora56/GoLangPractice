package main

import (
    "database/sql"
    "log"
    "fmt"

    "github.com/go-gorp/gorp"
    _ "github.com/mattn/go-sqlite3"
)

type Person struct {
    FirstName string `db:"first_name"`
    LastName  string `db:"last_name"`
    Email     string `db:"email"`
}

func main() {
    db, err := sql.Open("sqlite3", "test.db")
    if err != nil{
        log.Fatal(err)
    }
    dbmap := &gorp.DbMap{Db: db, Dialect:gorp.SqliteDialect{}}
    dbmap.AddTableWithName(Person{}, "person").SetKeys(false, "email")
    err = dbmap.CreateTablesIfNotExists()
    if err != nil{
        log.Fatalf("Can't create table: %v", err)
    }

    err = dbmap.Insert(&Person{"Momo", "Asakura", "mocho@musicrayn.com"})
    if err != nil{
        log.Fatalf("Insert error: %v", err)
    }

    var res []Person
    _, err = dbmap.Select(&res, "SELECT * FROM person")
    if err != nil{
        log.Fatalf("Select error: %v", err)
    }
    fmt.Println(res)    
}
