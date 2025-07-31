package handler

import (
	"net/http"
	"qr-code-generator/pkg/routes"
)

var router http.Handler

func init() {
	// initialize your mux router using existing routes
	router = routes.SetupRoutes()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
