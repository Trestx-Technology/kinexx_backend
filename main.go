package main

import (
	"kinexx_backend/pkg/router"
	"log"
	"net/http"

	trestCommon "github.com/Trestx-technology/trestx-common-go-lib"
	"github.com/rs/cors"
)

// setupGlobalMiddleware will setup CORS
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.AllowAll().Handler
	return handleCORS(handler)
}

// @title Kinexx API's
// @version 1.0
// @description This is a sample serice for managing kinexx
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:6019
// @BasePath /api/v1
func main() {

	trestCommon.LoadConfig()
	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(":6019", setupGlobalMiddleware(router)))
}
