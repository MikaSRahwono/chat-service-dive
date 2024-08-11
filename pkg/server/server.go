// server.go
package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	// Define routes
	return r
}

func StartServer(port string) {
	r := NewRouter()
	handler := cors.Default().Handler(r)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
