package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main(){
     router:= mux.NewRouter()
     router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
     	fmt.Fprint(w, "Hello world !!")

	 })

	http.ListenAndServe(":8080",router)


}
