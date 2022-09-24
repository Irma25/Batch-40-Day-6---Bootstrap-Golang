package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	// (req,res)
	route.HandleFunc("/", HandlerHello).Methods("GET")

	route.HandleFunc("/Contact", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("My Contact"))
	}).Methods("GET")

	fmt.Println("Server Running On Port 5000")
	http.ListenAndServe("localhost:5000", route)
}

func HandlerHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Duniaku"))
}
