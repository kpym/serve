package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// the version will be set by goreleaser based on the git tag
var version string = "dev"

// try logs the error and a message if needed
func try(e error, msg string) {
	if e != nil {
		log.Println(msg)
		log.Println(e)
	}
}

// mainEnd is the last function executed in this program.
func mainEnd() {
	// in case of error return status is 1
	if r := recover(); r != nil {
		os.Exit(1)
	}
	// the normal return status is 0
	os.Exit(0)
}

// If we terminate with Ctrl/Cmd-C we call mainEnd()
func catchCtrlC() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("server interrupted")
		mainEnd()
	}()
}

// SetPipeHanler sets the handler function in case of piped data
func SetPipeHanler(stdin []byte) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			log.Printf("serve from stdin")
			w.Write(stdin)
		} else {
			log.Printf("serve static file: " + r.URL.Path)
			http.FileServer(http.Dir(".")).ServeHTTP(w, r)
		}
	})
}

func main() {
	// error handling
	defer mainEnd()
	// interrupt handling
	catchCtrlC()

	// find the first available port after 8080
	var hostport string
	for i := 8080; i < 8180; i++ {
		hostport = fmt.Sprintf("localhost:%d", i)
		if ln, err := net.Listen("tcp", hostport); err == nil {
			ln.Close()
			break
		}
	}

	var server http.Handler
	// check if there is a piped data to serve
	fi, err := os.Stdin.Stat()
	if err == nil && (fi.Mode()&os.ModeNamedPipe != 0) {
		// if yes, serve the piped data to at "/"
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal("Error reading from stdin:", err)
		}
		log.Printf("serve [%s]: start serving the piped data to %s.", version, hostport)
		SetPipeHanler(stdin)
		server = nil
	} else {
		// if no, serve directly the current folder
		log.Printf("serve [%s]: start serving the current folder to %s.", version, hostport)
		server = http.FileServer(http.Dir("."))
	}

	// Try to open the browser before to start serving
	try(openbrowser("http://"+hostport), "Can't open the web browser.")
	log.Fatal(http.ListenAndServe(hostport, server))
}
