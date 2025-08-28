package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Handler for /process
func processHandler(w http.ResponseWriter, r *http.Request) {
	// Create a context with timeout (3 seconds)
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	// Channel to simulate work
	done := make(chan string, 1)

	// Start a goroutine that does some "work"
	go func() {
		time.Sleep(5 * time.Second) // pretend heavy work
		done <- "Work completed!"
	}()

	// Select: wait for work OR cancellation/timeout
	select {
	case result := <-done:
		fmt.Fprintln(w, result)
	case <-ctx.Done():
		// context is canceled (timeout or client disconnected)
		http.Error(w, "Request canceled or timed out", http.StatusRequestTimeout)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/process", processHandler)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
