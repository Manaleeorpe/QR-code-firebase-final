package routes

import (
	"log"
	"qr-code-generator/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterQRCodeGeneratorFirebasestoreRoutes = func(router *mux.Router) {
	log.Println("routes")
	router.HandleFunc("/firebase/qrcode/generate", controllers.QrcodeGenerateFirebase).Methods("POST")
	router.HandleFunc("/firebase/qrcode/{token}", controllers.ValidateQRCodeFirebase).Methods("GET")

}
