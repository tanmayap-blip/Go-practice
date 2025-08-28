package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// -------------------- MODEL --------------------
// User model with JSON tags (so it becomes {"id": 1, "name": "Alice"})
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// -------------------- STORAGE --------------------
var users = map[int]User{} // In-memory map to store users
var nextID = 1             // Auto-increment ID counter

// -------------------- MIDDLEWARE --------------------
// This middleware runs ONLY for the "getUser" handler
func getUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("📌 Middleware: Someone is trying to fetch a user by ID")
		// Pass the request to the actual handler
		next.ServeHTTP(w, r)
	})
}

// -------------------- HANDLERS --------------------

// Create user (POST /users)
func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode request body into User struct
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	// Assign ID and save in map
	u.ID = nextID
	nextID++
	users[u.ID] = u

	// Respond with created user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

// Get all users (GET /users)
func getUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Collect users into a slice
	list := []User{}
	for _, u := range users {
		list = append(list, u)
	}

	// Return JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

// Get user by ID (GET /users/{id})
func getUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL
	idStr := r.URL.Path[len("/users/"):]
	if idStr == "" {
		http.Error(w, "missing id in path", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	// Find user
	u, ok := users[id]
	if !ok {
		http.NotFound(w, r)
		return
	}

	// Return user as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

// Update user (PUT /users/{id})
func updateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/users/"):]
	if idStr == "" {
		http.Error(w, "missing id in path", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	// Check if user exists
	u, ok := users[id]
	if !ok {
		http.NotFound(w, r)
		return
	}

	// Decode request body
	var updatedData User
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	// Update fields
	u.Name = updatedData.Name
	users[id] = u

	// Respond with updated user
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

// Delete user (DELETE /users/{id})
func deleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/users/"):]
	if idStr == "" {
		http.Error(w, "missing id in path", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	// Remove user
	delete(users, id)

	// Respond with plain text
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "Deleted user %d", id)
}

// -------------------- MAIN --------------------
func main() {
	mux := http.NewServeMux()

	// Root
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("Welcome to the page"))
	})

	// /users (list or create)
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			createUser(w, r)
		} else if r.Method == http.MethodGet {
			getUsers(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// /users/{id} (get, update, delete)
	// Wrap only the GET handler with middleware
	mux.Handle("/users/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// Apply middleware here
			getUserMiddleware(http.HandlerFunc(getUser)).ServeHTTP(w, r)
		} else if r.Method == http.MethodPut {
			updateUser(w, r)
		} else if r.Method == http.MethodDelete {
			deleteUser(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}))

	fmt.Println("🚀 Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Server error:", err)
	}
}
