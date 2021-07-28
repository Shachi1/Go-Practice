package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome!</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "Email me at: <a href=\"mailtoabc@gmail.com\">abc@gmail.com</a>")
	} else {
		w.WriteHeader(404)
		fmt.Fprint(w, "<p>please email us if you keep being sent to the invalid page</p>")
	}
}
func main() {

	// var name string
	// fmt.Println("What is your name?")
	// // fmt.Scanf("%s", &name)
	// fmt.Println(name)

	mux := &http.ServeMux{}
	mux.HandleFunc("/", handlerFunc) 
	//http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", mux)
	// http.ListenAndServe(":3000", nil)

}
