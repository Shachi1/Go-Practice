package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

// func handlerFunc(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	if r.URL.Path == "/" {
// 		fmt.Fprint(w, "<h1>Welcome world!</h1>")
// 	} else if r.URL.Path == "/contact" || r.URL.Path == "/contact/"{
// 		fmt.Fprint(w, "Email me at: <a href=\"mailtoabc@gmail.com\">abc@gmail.com</a>")
// 	} else {
// 		w.WriteHeader(404)
// 		fmt.Fprint(w, "<p>please email us if you keep being sent to the invalid page</p>")
// 	}
// }

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome FAQ!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Email me at: <a href=\"mailtoabc@gmail.com\">abc@gmail.com</a>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(404)
	fmt.Fprint(w, "<p>please email us if you keep being sent to the invalid page</p>")
}

func main() {

	r := mux.NewRouter()
	// r.HandleFunc("/", handlerFunc)
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/notFound", notFound)
	http.ListenAndServe(":3000", r)
	// mux := &http.ServeMux{}
	// mux.HandleFunc("/", handlerFunc)
	//http.HandleFunc("/", handlerFunc)
	// http.ListenAndServe(":3000", mux)
	// http.ListenAndServe(":3000", nil)

}
