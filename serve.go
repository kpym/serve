package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	// find the first available port after 8080
	var hostport string
	for i := 8080; i < 8180; i++ {
		hostport = fmt.Sprintf("localhost:%d", i)
		if ln, err := net.Listen("tcp", hostport); err == nil {
			ln.Close()
			break
		}
	}
	// start serving the local folder
	log.Printf("Start serving the current folder to %s.", hostport)
	openbrowser("http://" + hostport)
	check(http.ListenAndServe(hostport, http.FileServer(http.Dir("."))))
}
