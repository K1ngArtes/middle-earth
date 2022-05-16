package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	if err != nil {
		log.Fatal("An error occurred when handling the request")
	}
}

func main() {
	// Right now default server mux is used
	//loggerMiddle := NewLogger(handler)
	http.HandleFunc("/bruh/lol", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "BRUH LOL!")
		if err != nil {
			log.Fatal("An error occurred when handling the BRUH LOL request")
		}
	})
	http.HandleFunc("/", handler)
	http.HandleFunc("/bruh", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "BRUH!")
		if err != nil {
			log.Fatal("An error occurred when handling the BRUH request")
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//Logger is a middleware handler that does request logging
type Logger struct {
	handler http.Handler
}

//ServeHTTP handles the request by passing it to the real
//handler and logging the request details
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}

//NewLogger constructs a new Logger middleware handler
func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}
