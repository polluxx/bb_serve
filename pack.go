package main

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

type Datatype struct {
	id		int
	name	string
	typed	int
	path	string
}

func main() {
	db, err := sql.Open("mysql", "root:@/bb")
    if err != nil {
        panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
    }
    defer db.Close()
	
	// Prepare statement for inserting data
    stmtIns, err := db.Prepare("INSERT INTO blocks VALUES( ?, ?, ?, ? )")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	stmtOut, err := db.Prepare("SELECT id,name,type,path FROM blocks")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
	var id, typed int
	var name, path string
    // Query the square-number of 13
    err = stmtOut.QueryRow().Scan(&id, &name, &typed, &path) // WHERE number = 13
    if err != nil {
		log.Fatal(err);
        //panic(err.Error()) // proper error handling instead of panic in your app
    }
    fmt.Print(id, name, typed, path)
	
    defer stmtOut.Close()
}