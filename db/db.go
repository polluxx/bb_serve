package db

import (
    "database/sql"
    //"fmt"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

type Datatype struct {
	id		int
	name	string
	typed	int
	path	string
}

func Connect(credentials string) *sql.DB {
	database, err := sql.Open("mysql", credentials)
	
    if err != nil {
        log.Fatal(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
    }
    defer database.Close()
	
	return database
}

func Insert(credents string, query string, insert []string) error{
	database, err := sql.Open("mysql", credents)
	
    if err != nil {
        log.Fatal(err)  
		return err
    }
    defer database.Close()
	
	// Prepare statement for inserting data
    stmtIns, err := database.Prepare(query)
    if err != nil {
		return err
    }
    defer stmtIns.Close() // Close the statement when we leave main() / the program terminates	
	
	_, err = stmtIns.Exec(nil, insert[0], insert[1])
	if (err != nil) {
		return err
	}
	
	return err
}

func QueryRow(credents string, query string) (map[int][]string, error) {	
	var result map[int][]string
	var err error
	
	database, err := sql.Open("mysql", credents)
	
    if err != nil {
        log.Fatal(err)  
		return result, err
    }
    defer database.Close()

	stmtOut, err := database.Query(query)
    if err != nil {
		log.Fatal(err)
        return result, err
    }
	
	columns, err := stmtOut.Columns()
	if(err != nil) {
		log.Fatal(err)
        return result, err		
	}
	
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
    for i := range values {
        scanArgs[i] = &values[i]
    }
	
	//fmt.Printf("%v", values)
	// Query data
    var index int = 0
	result = make(map[int][]string)
	
	for stmtOut.Next() {
		
		err = stmtOut.Scan(scanArgs...)
		if err != nil {
			log.Fatal(err)
			return result, err
		}
		
		resultItem := make([]string, len(values))
		
		for i, column := range values {
			if (column == nil) {
				resultItem[i] = "NULL"
			} else {
				resultItem[i] = string(column)
			}
		}
		
		result[index] = resultItem
		index++
	}
	defer stmtOut.Close()
	return result, err
}