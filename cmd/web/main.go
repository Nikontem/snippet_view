package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("GET	/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET	/{$}", home)
	mux.HandleFunc("GET	/snippet/view/{id}", snippetView)
	mux.HandleFunc("POST	/snippet/create", snippetCreatePost)

	log.Print("Starting server on :8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func snippetCreatePost(w http.ResponseWriter, request *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet"))
}
