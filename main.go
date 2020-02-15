package main

import (
	"encoding/json"
	"hello/fizzbuzz"
	"hello/oscar"
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
	r.HandleFunc("/oscarmale", oscarmalHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
func fizzbuzzHandler(w http.ResponseWriter, req *http.Request) {
	// vars := mux.Vars(req)
	// numberString := vars["number"]
	n, _ := strconv.Atoi(mux.Vars(req)["number"])
	io.WriteString(w, fizzbuzz.Say(n))
}

func oscarmalHandler(w http.ResponseWriter, req *http.Request) {
	m := oscar.ActorWhoGotMoreThanOne("./oscar/oscar_age_male.csv")

	w.Header().Set("Content-type", "text/json")
	json.NewEncoder(w).Encode(&m)
}

// func add(a int, b int) int {
// 	return a + b
// }
// func remove(a, b int) int {
// 	return a - b
// }

// mecho
// go datebase-sql
