package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", echo)
	if len(os.Args) < 2 {
		log.Fatalln("Please provide addr:port to bind to as a first argument")
	}
	log.Fatal(http.ListenAndServe(os.Args[1], nil))
}

func echo(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	defer body.Close()
	total, err := io.Copy(w, body)
	if err != nil {
		log.Printf("cannot write response: %v\n", err)
	}
	if total == 0 {
		_, err = w.Write([]byte("This is echo service, try posting something\n"))
		if err != nil {
			log.Printf("cannot write response: %v\n", err)
		}
	}
}
