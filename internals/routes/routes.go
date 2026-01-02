package routes

import (
	"fmt"
	"net/http"

	"github.com/JDarwind/go-skeleton-starter/internals/middlewares"
	"github.com/JDarwind/go-skeleton-starter/internals/requests"
	"github.com/JDarwind/go-skeleton-starter/pkg/network/httpkit"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	data, errors := httpkit.ValidateRequest(r, &requests.HelloRequest{})
	if errors != nil {
		httpkit.NewResponse(w, r).
			Status(http.StatusUnprocessableEntity).
			Error(errors)
		return
	}

	req := data.(requests.HelloRequest)

	ret := fmt.Sprintf("Hello %s", req.Name)

	httpkit.NewResponse(w, r).Success(ret)
}

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello-world/", helloWorld)

	chain := httpkit.Chain(
		mux,
		middlewares.Logging,
	)

	return chain
}
