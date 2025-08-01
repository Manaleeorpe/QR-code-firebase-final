package main

import (
	"log"
	"net/http"

	"qr-code-generator/pkg/config"
	"qr-code-generator/pkg/routes"

	"github.com/rs/cors"
)

func main() {
	/* Initialize database connection SQL
	config.Connect()
	db := config.GetDB()
	models.SetDB(db)
	db.AutoMigrate(&models.QRCode{})*/
	log.Println("something started")
	config.ConnectFirebase()

	/* Use Gorilla Mux
	router := mux.NewRouter()
	routes.RegisterQRCodeGeneratorstoreRoutes(router)
	routes.RegisterQRCodeGeneratorFirebasestoreRoutes(router)*/
	router := routes.SetupRoutes()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // React dev server
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	})
	handler := c.Handler(router)
	
	
	log.Println("Starting server on localhost:8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", handler)) // Use mux router here

}
