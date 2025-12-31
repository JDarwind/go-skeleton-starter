package routes

import (
	"net/http"
	"fmt"
	"github.com/JDarwind/go-skeleton-starter/pkg/network"
	"github.com/JDarwind/go-skeleton-starter/internals/middlewares"
)

func helloWorld(w http.ResponseWriter,r *http.Request){
	 fmt.Fprintf(w, "hello")
}


func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello-world/", helloWorld)
	
	chain:= network.Chain(
		mux,
		middlewares.Logging,
	)
	

	return chain;
}