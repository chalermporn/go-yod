package main

import (
	"hello/fizzbuzz"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	// Hello world, the web server

	// helloHandler := func(w http.ResponseWriter, req *http.Request) {
	// 	io.WriteString(w, "Hello, world!\n")
	// }
	r := mux.NewRouter()

	r.HandleFunc("/fizzbuzz/{number}", fizzbuzzHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
func fizzbuzzHandler(w http.ResponseWriter, req *http.Request) {
	// vars := mux.Vars(req)
	// numberString := vars["number"]
	n, _ := strconv.Atoi(mux.Vars(req)["number"])
	io.WriteString(w, fizzbuzz.Say(n))
}

// func add(a int, b int) int {
// 	return a + b
// }
// func remove(a, b int) int {
// 	return a - b
// }

// mecho
// go datebase-sql
