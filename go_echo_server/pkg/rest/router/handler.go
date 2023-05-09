package router

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


func InitializeHandlers() *chi.Mux {
	router := chi.NewRouter()
	// A good base middleware stack
	// Full list of middleware: https://go-chi.io/#/pages/middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger) // Logged should be before Recoverer
	router.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))

	// these handlers are in handleInvalid.go
	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)

	// switch on the method
	router.Route("/", func(r chi.Router) {
		r.Get("/", HandleReading)
		r.Get("/{path}", HandleReading)
		r.Post("/", HandleAdding)
		r.Post("/{path}", HandleAdding)
		r.Put("/{path}", HandleUpdating)
		r.Patch("/{path}", HandleUpdating)
		r.Delete("/{path}", HandleDeleting)
	})

	return router
}
