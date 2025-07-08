package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleReq(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Procedure Hello world everytime i learn something new")
}

func main() {
	// register the handle function to the DefaultServeMux
	http.HandleFunc("/", handleReq)
	//serve on 8080, set handler to nil so that it will use DefaultServeMux
	log.Fatal(http.ListenAndServe(":8080", nil))
}
