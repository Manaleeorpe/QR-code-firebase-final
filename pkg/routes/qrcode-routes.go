package routes

import (
	"qr-code-generator/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterQRCodeGeneratorstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/generate", controllers.QrcodeGenerate).Methods("POST")
	router.HandleFunc("/qrcode/{token}", controllers.ValidateQRCode).Methods("GET")

}
