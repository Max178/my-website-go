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
    // Serve static files from the images directory
    fs := http.FileServer(http.Dir("images"))
    http.Handle("/images/", http.StripPrefix("/images/", fs))
    
    // Register handlers
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/home", homeHandler)
    http.HandleFunc("/blog", blogHandler)
    
    // Start server - bind to all interfaces (0.0.0.0)
    port := ":80"
    fmt.Printf("Server running on port %s\n", port)
    log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
}

