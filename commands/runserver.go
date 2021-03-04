package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main()  {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to this life-changing API.\nIt's the best API, it's true, all other API's are fake")
	})

	fmt.Println("Server listening!")
	http.ListenAndServe(":80", r)
}
