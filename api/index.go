package handler

import (
	"net/http"

	"qr-code-generator/pkg/config"
	"qr-code-generator/pkg/routes"

	"github.com/gorilla/mux"
)

var router *mux.Router

func init() {
	config.ConnectFirebase()
	router = mux.NewRouter()
	routes.RegisterQRCodeGeneratorstoreRoutes(router)
	routes.RegisterQRCodeGeneratorFirebasestoreRoutes(router)
}

// Handler is Vercel entry point
func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
