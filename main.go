package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// FileServer
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(writer, "404 not found", http.StatusNotFound)
		return
	}
	if request.Method != "GET" {
		http.Error(writer, "method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(writer, "hello!")
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		fmt.Fprintf(writer, "parseForm() err: %v", err)
		return
	}
	fmt.Fprintf(writer, "POST request succesfull")
	name := request.FormValue("name")
	address := request.FormValue("address")
	fmt.Fprintf(writer, "Name = %s\n", name)
	fmt.Fprintf(writer, "Address = %s\n", address)

}
