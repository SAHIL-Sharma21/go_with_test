package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// fmt.Println("Depenedency injection using test")
	// Greet(os.Stdout, "Sharma")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeHandler)))
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Sahil")
}
