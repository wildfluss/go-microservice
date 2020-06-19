package main 

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request from %v", r.RemoteAddr)
	hostname, err := os.Hostname()
	if err != nil {
		// TODO logging 
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}
	fmt.Fprintf(w, "You've hit %s\n", hostname) // 200
}