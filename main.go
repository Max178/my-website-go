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
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/blog", blogHandler)
	
	// Start server - bind to all interfaces (0.0.0.0)
	port := ":80"
	fmt.Printf("Server running on port %s\n", port)
	fmt.Println("Endpoints:")
	fmt.Println("  GET /home")
	fmt.Println("  GET /blog")
	fmt.Println("Access via: http://:8080")
	
	log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
}

