package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"title":   "About",
		"content": "This is the About section of my portfolio",
	})
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to return the projects data as JSON
}

func contactsHandler(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to return the contact data as JSON
}

func skillsHandler(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to return the skills data as JSON
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	http.HandleFunc("/api/about", aboutHandler)
	http.HandleFunc("/api/projects", projectsHandler)
	http.HandleFunc("/api/contact", contactsHandler)
	http.HandleFunc("/api/skills", skillsHandler)

	fmt.Println("Server listening on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
