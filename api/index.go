package handler

import (
    "net/http"

    "github.com/gorilla/mux"
    "qr-code-generator/pkg/routes"
)

var router *mux.Router

func init() {
    router = mux.NewRouter()
    routes.RegisterQRCodeGeneratorstoreRoutes(router)
    routes.RegisterQRCodeGeneratorFirebasestoreRoutes(router)
}

// Handler is Vercel entry point
func Handler(w http.ResponseWriter, r *http.Request) {
    router.ServeHTTP(w, r)
}
