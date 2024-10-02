package main

// Imports
import (
	"fmt"
	"net/http"
)

// Handlers
func HandleHello(w http.ResponseWriter, r *http.Request) { 
	w.Write([]byte("Meow"))
}

// Entry
func main() {
	// Server 
	server := http.NewServeMux(); 
	// Routes 
	http.HandleFunc("/hello",  HandleHello)
	
	err := http.ListenAndServe(":7777", server)
	
	if err == nil { 
		fmt.Println("Error while opening the server!")
	}
}