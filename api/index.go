package handler

import (
	"fmt"
	"log"
	"net/http"
)

//var router http.Handler
var x int
func init() {
	// initialize your mux router using existing routes
	//router = routes.SetupRoutes()
	log.Println("Router initialized")
	x = 12
}

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Router initialized from handler")
	fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
	fmt.Fprintf(w, "<h1>Hello from Go! x = %d</h1>", x)
}

/*func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}*/

/*package handler

import (
  "fmt"
  "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}*/
