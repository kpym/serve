package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func try(e error, msg string) {
	if e != nil {
		log.Println(msg)
		log.Println(e)
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
	try(openbrowser("http://"+hostport), "Can't open the web browser.")
	log.Fatal(http.ListenAndServe(hostport, http.FileServer(http.Dir("."))))
}
