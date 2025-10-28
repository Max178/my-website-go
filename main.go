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
    ip := r.RemoteAddr
    if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
        ip = forwarded
    }
    log.Printf("Home page accessed from IP: %s", ip)
}

// blogHandler handles requests to /blog
func blogHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "blog.html")
    ip := r.RemoteAddr
    if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
        ip = forwarded
    }
}

// contactHandler handles requests to /contact
func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "contact.html")
    ip := r.RemoteAddr
    if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
        ip = forwarded
    }
}

// projectHandler handles requests to /project
func projectHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	http.ServeFile(w, r, "projects.html")
    log.Println("Project page accessed")
    ip := r.RemoteAddr
    if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
        ip = forwarded
    }
}

// aboutHandler handles requests to /about
func aboutHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    http.ServeFile(w, r, "about.html")
    ip := r.RemoteAddr
    if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
        ip = forwarded
    }
}

func main() {
    // Serve static files from the images directory
    fs := http.FileServer(http.Dir("images"))
    http.Handle("/images/", http.StripPrefix("/images/", fs))
    
    // Register handlers
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/home", homeHandler)
    http.HandleFunc("/blog", blogHandler)
    http.HandleFunc("/contact", contactHandler)
    http.HandleFunc("/projects", projectHandler)
    http.HandleFunc("/about", aboutHandler)
    
    // Start server - bind to all interfaces (0.0.0.0)
    port := ":80"
    fmt.Printf("Server running on port %s\n", port)
    // For HTTP only
    log.Fatal(http.ListenAndServe(":80", nil))
    // For HTTPS only (requires SSL certificates)
}
