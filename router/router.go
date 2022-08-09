package router

import (
	"movies-rest-api/middlewares"
	"movies-rest-api/services"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

func Router() *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.AllowAll().Handler)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/api", func(r chi.Router) {

		r.Route("/movies", func(r chi.Router) {
			r.Use(middlewares.Auth)
			r.Get("/", services.GetAllMovies)
			r.Get("/{id}", services.GetMovieByID)
			r.Get("/search", services.SearchMovies)
		})

		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", services.UserLogin)
			r.Post("/register", services.UserSignup)
			
			r.Route("/", func(r chi.Router) {
				r.Use(middlewares.Auth)
				r.Post("/logout", services.Logout)
				r.Post("/verify", services.Verify)
			})
		})
	})
	return router
}
