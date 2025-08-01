package routes

import (
    "net/http"
    "qr-code-generator/pkg/controllers"

    "github.com/gorilla/mux"
)

// Base URL handler - responds with "hi"
func SayHiHandler(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}

// Register all your Firebase QR code related routes here
var RegisterQRCodeGeneratorFirebasestoreRoutes = func(router *mux.Router) {
    router.HandleFunc("/", SayHiHandler).Methods("GET") // base URL route
    router.HandleFunc("/firebase/qrcode/generate", controllers.QrcodeGenerateFirebase).Methods("POST")
    router.HandleFunc("/firebase/qrcode/{token}", controllers.ValidateQRCodeFirebase).Methods("GET")
}

// If you have other QR code routes, you can register them he

// SetupRoutes returns a mux.Router with all routes registered
func SetupRoutes() *mux.Router {
    router := mux.NewRouter()
    RegisterQRCodeGeneratorstoreRoutes(router)
    RegisterQRCodeGeneratorFirebasestoreRoutes(router)
    return router
}
