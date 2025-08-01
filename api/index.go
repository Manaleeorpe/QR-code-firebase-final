import (
	"log"
	"net/http"
)

var router http.Handler

func init() {
	// initialize your mux router using existing routes
	//router = routes.SetupRoutes()
	log.Println("Router initialized")
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
