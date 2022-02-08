package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./src")))

	http.HandleFunc("/counter", incrementCounter)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
