package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)

	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	var listenOn = ":" + port

	http.HandleFunc("/", handler)
	log.Printf("Listening on %v", listenOn)
	var err error = http.ListenAndServe(listenOn, nil)
	log.Fatalf("Server exited with %v", err)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var body, err = ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Cannot read body: %v", err)
		return
	}
	log.Printf("Got body `%s` at %v in request %+v ", body, r.RequestURI, r)
	w.Write([]byte("OK"))
}
