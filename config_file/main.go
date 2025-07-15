package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

const cfgPath = "/config/demo.yaml"

var config struct {
	Greeting string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s, yello\n", config.Greeting)
}

func main() {
	cfgData, err := os.ReadFile(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	if err = yaml.Unmarshal(cfgData, &config); err != nil {
		log.Fatalf("failed to parse config file %q: %v", cfgPath, err)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
