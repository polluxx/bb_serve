package main

import (
    "github.com/polluxx/bb_serve/db"
	//"net/http"
	"fmt"
)

type Datatype struct {
	id		int
	name	string
	typed	int
	path	string
}

func main() {
	var credentials string = "root:@/bb"
	//database := db.Connect(credentials)
	
	insert := make([]string, 2)
	insert[0] = "0"
	insert[1] = "name"
	
	queryInsert := fmt.Sprintf("INSERT INTO %s VALUES( ?, ?, ?)", "types")
	
	err := db.Insert(credentials, queryInsert, insert)
	if (err != nil) {
		fmt.Printf("%s", err.Error())
	}
	
	result, err := db.QueryRow(credentials, "select * from types");
	if (err != nil) {
		fmt.Printf("%s", err.Error())
	}
	
	fmt.Printf("%v", result)
}