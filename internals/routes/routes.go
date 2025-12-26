package routes

import (
	"net/http"
	"fmt"
)

func helloWorld(w http.ResponseWriter,r *http.Request){
	 fmt.Fprintf(w, "hello")
}


func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello-world/", helloWorld)
	return mux;
}