package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

const dbFile = "dbFile"

var db tinyDB

// Simple data structure to persist the counter
type tinyDB struct {
	FileName string
	Mx       sync.Mutex
	Counter  int
}

func incrementHandler(db *tinyDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db.Mx.Lock()
		defer db.Mx.Unlock()
		inputBytes, err := os.ReadFile(db.FileName)
		check(err)
		inputString := string(inputBytes)
		counter, err := strconv.Atoi(inputString)
		check(err)
		counter++
		outputString := strconv.Itoa(counter)
		fmt.Fprintf(w, outputString)
		bytes := []byte(outputString)
		writeErr := os.WriteFile(db.FileName, bytes, 0600)
		check(writeErr)
	}
}

// "Error handling"
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func init() {
	_, err := os.Stat(dbFile)
	if err != nil {
		f, err := os.Create(dbFile)
		f.WriteString("0")
		check(err)
		f.Close()
	}
}

func main() {

	db := tinyDB{FileName: dbFile, Mx: sync.Mutex{}, Counter: 0}

	http.Handle("/", http.FileServer(http.Dir("./src")))

	http.HandleFunc("/counter", incrementHandler(&db))

	log.Fatal(http.ListenAndServe(":8081", nil))
}
