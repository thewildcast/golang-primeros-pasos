package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Backend struct {
	port string
}

func (b *Backend) LBHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Path: %s", html.EscapeString(r.URL.Path))
	log.Printf("Method: %s", html.EscapeString(r.Method))
	log.Printf("Host: %s", html.EscapeString(r.Host))
	log.Printf("Protocol: %s", html.EscapeString(r.Proto))
	body, _ := ioutil.ReadAll(r.Body)
	log.Printf("Body: %v", body)
	log.Printf("Headers: %v", r.Header)
	w.Header().Add("brunoli", "de wilde")
	fmt.Fprintf(w, "Soy el backend del puerto: "+b.port)

	log.Println("===========================================================================")

}

func main() {
	port := os.Args[1]
	log.Printf("Listening on port %s", port)
	b := Backend{port: port}
	http.HandleFunc("/", b.LBHandler)
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}
