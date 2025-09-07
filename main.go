package main

import (
	"fmt"
	"log"
	"net/http"
)

// homeHandler handles requests to /home
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "home.html")
}

// blogHandler handles requests to /blog
func blogHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "blog.html")
}

func main() {
	// Register handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/blog", blogHandler)
	
	// Start server
	port := ":8080"
	fmt.Printf("Server running on http://localhost%s\n", port)
	fmt.Println("Endpoints:")
	fmt.Println("  GET /")
	fmt.Println("  GET /blog")
	
	log.Fatal(http.ListenAndServe(port, nil))
}