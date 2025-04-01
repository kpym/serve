package main

import (
	_ "embed"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// the version will be set by goreleaser based on the git tag
var version string = "dev"

//go:embed serve.ico
var favIcon []byte

// get the command line arguments
var (
	port string
	path string
)

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

// serveAtRoot returns a handler that serves the given data at the root path.
// all other paths are served by h
func serveAtRoot(data []byte, h http.Handler) http.Handler {
	// if data is empty, return the handler as is
	if len(data) == 0 {
		return h
	}
	// if data is not empty, serve it at root
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" || r.URL.Path == "" {
			log.Printf("Serve piped data at root.")
			w.Write(data)
		} else {
			h.ServeHTTP(w, r)
		}
	})
}

// logHandler wraps the given handler with a logging middleware that logs the request path.
func logHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s", r.URL.Path)
		h.ServeHTTP(w, r)
	})
}

// favHandler wraps the given handler with a middleware that serves the favicon.
func favHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/favicon.ico" {
			// Serve the favicon
			w.Header().Set("Cache-Control", "max-age=86400") // 86400 s = 1 day
			w.Header().Set("Expires", time.Now().Add(24*time.Hour).UTC().Format(http.TimeFormat))
			w.Write(favIcon)
			return
		}
		// Delegate other requests to the original handler
		h.ServeHTTP(w, r)
	})
}

func init() {
	// set help message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `serve (version: %s): Serve a directory or stdin as a webserver.
Usage: serve [options]

Options:

`, version)
		flag.PrintDefaults()
	}
	// set command line arguments
	flag.StringVar(&port, "p", "", "port to listen on")
	flag.StringVar(&path, "t", "", "serve at the given path (examples: 'foo' or '/foo/bar/')")

	// init the log
	log.SetFlags(log.LstdFlags | log.Lmsgprefix)
	log.SetPrefix("[serve] ")
}

func main() {
	// error handling
	defer mainEnd()
	// interrupt handling
	catchCtrlC()

	// parse the command line arguments
	flag.Parse()

	// set host port to 'localhost:port'
	var hostport string
	if port != "" {
		// if port is given, use it
		hostport = "localhost:" + port
	} else {
		// else find the first available port after 8080
		for i := 8080; i < 8180; i++ {
			hostport = fmt.Sprintf("localhost:%d", i)
			if ln, err := net.Listen("tcp", hostport); err == nil {
				ln.Close()
				break
			}
		}
	}

	// set the handler
	server := http.FileServer(http.Dir("."))
	// check if there is a piped data to serve
	fi, err := os.Stdin.Stat()
	if err == nil && (fi.Mode()&os.ModeNamedPipe != 0) {
		// if yes, serve the piped data at root
		stdin, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal("Error reading from stdin:", err)
		}
		log.Printf("Piped data is present.")
		server = serveAtRoot(stdin, server)
	}
	// check if there is a path to serve to
	if path != "" {
		// normalize path
		path = "/" + strings.Trim(path, "/") + "/"
		server = http.StripPrefix(path, server)
	}
	// log the request path
	server = logHandler(favHandler(server))

	// Try to open the browser before to start serving
	try(openbrowser("http://"+hostport+path), "Can't open the web browser.")
	log.Printf("Start serving the current folder at http://%s%s.", hostport, path)
	log.Fatal(http.ListenAndServe(hostport, server))
}
