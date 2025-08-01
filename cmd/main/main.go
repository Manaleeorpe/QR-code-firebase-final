package main

import (
	"log"
	"net/http"
	"os"

	"qr-code-generator/pkg/config"
	"qr-code-generator/pkg/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Connect Firebase
	config.ConnectFirebase()

	// Use Gorilla Mux
	router := mux.NewRouter()
	routes.RegisterQRCodeGeneratorstoreRoutes(router)
	routes.RegisterQRCodeGeneratorFirebasestoreRoutes(router)

	// CORS settings
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // React dev server
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	})
	handler := c.Handler(router)

	// Get port from environment variable or default to 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Starting server on 0.0.0.0:%s...\n", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, handler))
}
