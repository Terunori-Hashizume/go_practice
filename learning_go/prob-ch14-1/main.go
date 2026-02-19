package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func Middleware(ms int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx, cancel := context.WithTimeout(ctx, time.Duration(ms)*time.Millisecond)
			defer cancel()
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func run() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			time.Sleep(5 * time.Millisecond)
			if r.Context().Err() == context.DeadlineExceeded {
				fmt.Fprint(w, "Timeout")
				return
			}
			fmt.Fprint(w, "Hello, world!")
		}
	})

	// handler := Middleware(1)(mux) // こちらはタイムアウトする
	handler := Middleware(10)(mux) // こちらはタイムアウトしない
	http.ListenAndServe(":8080", handler)
}

func main() {
	run()
}
