package handler
 
import (
  "fmt"
  "net/http"
  "log"		
)
func init() {
	// initialize your mux router using existing routes
	//router = routes.SetupRoutes()
	log.Println("Router initialized")
}
 
func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}
