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

type Type struct {
	Id		null
	Name	string
	Own 	string
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
    if err != null {
		return err
    }
    defer stmtIns.Close() // Close the statement when we leave main() / the program terminates	
	
	typed := Type{nil, insert[0], insert[1]}
	
	_, err = stmtIns.Exec(&typed)
	if (err != nil) {
		return err
	}
	
	return err
}

func QueryRow(credents string, query string) ([]string, error) {	
	var result []string
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
	// Query data
    
	for stmtOut.Next() {
		err = stmtOut.Scan(&values)
		if err != nil {
			log.Fatal(err)
			return result, err
		}
		
		for i, column := range values {
			if (column == nil) {
				result[i] = "NULL"
			} else {
				result[i] = string(column)
			}
		}
	}
	defer stmtOut.Close()
	return result, err
}