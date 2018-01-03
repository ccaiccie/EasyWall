package main

import (
	"log"
	"net/http"
	// "encoding/json"
	"io"
	"net/http/httputil"
)

// type handler func(w http.ResponseWriter, r *http.Request)

// StartWebserver - Start the Listener for Pages, Assets and AJAX
func StartWebserver() {
	// Handle Assets Static
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handle Individual Pages
	http.HandleFunc("/", homeHandler)

	// Start the Server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// homeHandler - Handler for the Index Page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

	} else if r.Method == "GET" {
		// w.Header().Set("Content-Type", "application/json")
		result, _ := httputil.DumpRequest(r, true)
		Log(string(result))
		io.WriteString(w, string(result))
	}
}

// errorHandler - Handle Not Found or 403 Errors Correctly
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		http.Error(w, "404", http.StatusNotFound)
	}
}

// func HandleIndex(w http.ResponseWriter, r *http.Request) {
//     io.WriteString(w, "hello, world\n")
// }

// func HandlePost(w http.ResponseWriter, r *http.Request) {
//     r.ParseForm()
//     log.Println(r.PostForm)
//     io.WriteString(w, "post\n")
// }

// type Result struct {
//     FirstName string `json:"first"`
//     LastName  string `json:"last"`
// }

// func HandleJSON(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")
//     result, _ := json.Marshal(Result{"tee", "dub"})
//     io.WriteString(w, string(result))
// }
