package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// main funciton is en entry point for execution lik C & C++
func main() {
	r := mux.NewRouter()
	form := "<html><head></head><body>" +
		"<form action='' method='post'>" +
		"Name:<input type='text' name='name'>" +
		"</body></html>"
	func1 := func(w http.ResponseWriter, r *http.Request) {
		//		vars := mux.Vars(r)
		//		w.Write([]byte(vars["id"]))
		w.Write([]byte(form))
	}

	//get the form value through POST method
	func2 := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.FormValue("name")))
	}
	/*
		func2 := func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Funciton 2"))
		}

		handler := func(w http.ResponseWriter, r *http.Request) {
			if r.TLS == nil {
				func1(w, r)
			} else {
				func2(w, r)
			}
		}*/
	// work with path matches
	//	r.HandleFunc("/foo/{id:[0-9]+}", func1)
	//  Handle requests with Method Specifides
	r.HandleFunc("/", func1).Methods("GET")
	r.HandleFunc("/", func2).Methods("POST")
	http.Handle("/", r)
	go http.ListenAndServe(":4000", nil)

	//go http.ListenAndServeTLS(":4000", nil)
	fmt.Scanln()
}
