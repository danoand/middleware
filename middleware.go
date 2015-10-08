package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Middleware handler function that logs inbound request information
func loggingHandler(next http.Handler) http.Handler {
	// Define the logging code
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("RQST: [%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}

	// Return a handler function that wraps the logging code and the core handler function
	return http.HandlerFunc(fn)
}

// Middleware handler function that recovers from a panic in the underlying request handler (if it occurs)
func recoverHandler(next http.Handler) http.Handler {
	// Define a function that defers a function to recover from a panic
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// Middleware handler function that will validate that a request is a POST and return if it is not
func valPOST(next http.Handler) http.Handler {
	m := "POST"

	// Define a function that passes on a POST but returns for any other type of request
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Check the existing request method
		if r.Method == m {
			// Request method is a post, continue down the chain of middleware
			next.ServeHTTP(w, r)
		} else {
			// Request is NOT a post return
			log.Printf("INFO: Request from %v is not a %v.  Not continuing.\n", r.URL.String(), m)
			// Reply to the request with an Error
			http.Error(w, fmt.Sprintf("STATUS = Method [%v] not allowed.", r.Method), http.StatusMethodNotAllowed)
		}
	}

	return http.HandlerFunc(fn)
}

// Middleware handler function that will validate that a request is a GET and return if it is not
func valGET(next http.Handler) http.Handler {
	m := "GET"

	// Define a function that passes on a GET but returns for any other type of request
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Check the existing request method
		if r.Method == m {
			// Request method is a post, continue down the chain of middleware
			next.ServeHTTP(w, r)
		} else {
			// Request is NOT a post return
			log.Printf("INFO: Request from %v is not a %v.  Not continuing.\n", r.URL.String(), m)
			// Reply to the request with an Error
			http.Error(w, fmt.Sprintf("STATUS = Method [%v] not allowed.", r.Method), http.StatusMethodNotAllowed)
		}
	}

	return http.HandlerFunc(fn)
}
