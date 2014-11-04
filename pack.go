package main

import (
    "github.com/polluxx/bb_serve/db"
	"net/http"
	"fmt"
	"time"
	"encoding/json"
	"log"
	//"os"
	"regexp"
)

type Datatype struct {
	id		int
	name	string
	typed	int
	path	string
}

var Credentials string = "root:gtnhjdbx@/bb"

func main() {
    	http.HandleFunc("/select", mainHandler(selectHandler))
    	http.HandleFunc("/insert", mainHandler(insertHandler))

	s := &http.Server{
		Addr:           ":8090",
		//Handler:        Handle,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	
	log.Fatal(s.ListenAndServe())
}

func mainHandler (fn func(http.ResponseWriter, *http.Request, map[string]string)) http.HandlerFunc {
    var validPath = regexp.MustCompile("^/(insert|select)")
        return func(w http.ResponseWriter, r *http.Request) {
		mess := validPath.FindStringSubmatch(r.URL.Path)
		if mess == nil {
			http.NotFound(w,r);
			return
		}
		
		r.ParseForm();
		queryParams := make(map[string]string)
		for index, value := range r.Form {
			queryParams[index] = value[0];
		}
		
		fn(w, r, queryParams)
	}
}

func selectHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	
	result, err := db.QueryRow(Credentials, "select * from types");
	if (err != nil) {
		fmt.Printf("%s", err.Error())
	}
	
	Response(w,r, result)
	//fmt.Printf("%T", result)
}

func insertHandler(w http.ResponseWriter, r *http.Request, params map[string]string) {
	
	insert := make([]string, 2)
	queryInsert := fmt.Sprintf("INSERT INTO %s VALUES( ?, ?, ?)", "types")
	
	err := db.Insert(Credentials, queryInsert, insert)
	if (err != nil) {
		fmt.Printf("%s", err.Error())
	}
}

func Response(w http.ResponseWriter, r *http.Request, params map[int][]string) {
	
	resp := make(map[string][]string)
	
	for key, item := range params {
		resp[fmt.Sprintf("%d", key)] = item
	}

	jsn, err := json.Marshal(resp)
	
	if err != nil {
	       	http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "false")
	w.Header().Set("Access-Control-Allow-Headers", "accept, authorization")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	
	w.Write(jsn);
}
