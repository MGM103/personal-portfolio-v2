package app

import (
	"fmt"
	"net/http"
)

func NewServer(port string) *http.Server {
	mux := http.NewServeMux()

	RegisterRoutes(mux)

	return &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}
}
