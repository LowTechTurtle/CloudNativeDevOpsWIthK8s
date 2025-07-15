package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const cfgPath = "/secrets/password"

var secret []byte

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s, yello\n", secret)
}

func main() {
	var err error
	secret, err = os.ReadFile(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
