package routes

import (
	"github.com/go-chi/chi"
	"github.com/onmybob/aileg-server/controllers"
)

func EchoRoutes(r *chi.Mux){
	r.Get("/", controllers.Echo)
}