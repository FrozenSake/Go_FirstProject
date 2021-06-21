package main

import (
	"fmt"
	"log"
	"net/http"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func main() {

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		response := About()
		fmt.Fprint(w, response)
	})

	http.HandleFunc("/echo", echoHandler)
	http.HandleFunc("/invert", invertHandler)
	http.HandleFunc("/flatten", flattenHandler)
	http.HandleFunc("/sum", sumHandler)
	http.HandleFunc("/multiply", multiplyHandler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("Server failed to start. error %s", err.Error())
		return
	}
}

// Handler for echo
func echoHandler(w http.ResponseWriter, r *http.Request) {
	records, err := GetFileMatrix(r)
	if !ErrorHandler(w, err, "could not retrive file.") {
		return
	}

	response, err := Echo(records)
	if !ErrorHandler(w, err, "echo: ") {
		return
	}

	fmt.Fprint(w, response)
}

// Handler for invert
func invertHandler(w http.ResponseWriter, r *http.Request) {
	records, err := GetFileMatrix(r)
	if !ErrorHandler(w, err, "could not retrive file.") {
		return
	}

	response, err := Invert(records)
	if !ErrorHandler(w, err, "invert: ") {
		return
	}

	fmt.Fprint(w, response)
}

// Handler for flatten
func flattenHandler(w http.ResponseWriter, r *http.Request) {
	records, err := GetFileMatrix(r)
	if !ErrorHandler(w, err, "could not retrive file.") {
		return
	}

	response, err := Flatten(records)
	if !ErrorHandler(w, err, "flatten: ") {
		return
	}

	fmt.Fprint(w, response)
}

// Handler for sum
func sumHandler(w http.ResponseWriter, r *http.Request) {
	records, err := GetFileMatrix(r)
	if !ErrorHandler(w, err, "could not retrive file.") {
		return
	}

	response, err := Sum(records)
	if !ErrorHandler(w, err, "sum: ") {
		return
	}

	fmt.Fprint(w, response)
}

// Handler for multiply
func multiplyHandler(w http.ResponseWriter, r *http.Request) {
	records, err := GetFileMatrix(r)
	if !ErrorHandler(w, err, "could not retrive file.") {
		return
	}

	response, err := Multiply(records)
	if !ErrorHandler(w, err, "multiply: ") {
		return
	}

	fmt.Fprint(w, response)
}
