package app

import (
	"net/http"

	"github.com/mgm103/personal-portfolio-v2/internal/app/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", handlers.Home)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}
