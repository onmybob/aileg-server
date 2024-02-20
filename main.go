package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/onmybob/aileg-server/routes"
	"github.com/onmybob/aileg-server/utils"
)

var(	
	//stderrLogger = log.New(os.Stderr, "aileg-server: ", log.Llongfile)


)


func main() {
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	utils.Log.Info().Msg("232332322332")

	r := chi.NewRouter()
	SetMiddleware(r)
	routes.EchoRoutes(r)

    err :=http.ListenAndServe(":"+port, r)

	if err != nil {
		utils.Log.Fatal().Msg("232332322332")
	}
}

func SetMiddleware(r *chi.Mux){
	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentType("application/json","text/xml"))
	r.Use(middleware.RealIP)
}