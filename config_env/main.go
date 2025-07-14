package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handleReq(w http.ResponseWriter, r *http.Request) {
	turtle := os.Getenv("big-turtle")
	fmt.Fprintln(w, "hi ", turtle)
}

func main() {
	// register the handle function to the DefaultServeMux
	http.HandleFunc("/", handleReq)
	//serve on 8080, set handler to nil so that it will use DefaultServeMux
	log.Fatal(http.ListenAndServe(":8080", nil))
}
