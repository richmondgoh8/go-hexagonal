package main

import (
	"github.com/go-chi/chi/v5"
	svc "github.com/richmondgoh8/boilerplate/internal/core/services"
	handler "github.com/richmondgoh8/boilerplate/internal/handlers"
	"github.com/richmondgoh8/boilerplate/pkg/logger"
	custommiddleware "github.com/richmondgoh8/boilerplate/pkg/middleware"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/richmondgoh8/boilerplate/internal/platform/config"
	"github.com/richmondgoh8/boilerplate/internal/platform/db"
	repo "github.com/richmondgoh8/boilerplate/internal/repositories/postgres"
	apperror "github.com/richmondgoh8/boilerplate/static"

	_ "github.com/lib/pq"
)

func main() {
	config.InitReader()
	logger.Init()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal(apperror.EmptyPort)
	}

	r := chi.NewRouter()

	// Start of Middleware
	r.Use(custommiddleware.InjectTrackingID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// End of Middleware

	// Start of External Dependencies Instantiation
	localDB, err := db.Init()
	if err != nil {
		panic("error connecting to db")
	}
	// End of External Dependencies Instantiation

	// Start of Dependency Injection
	linkRepo := repo.NewPostgresInstance(localDB)
	linkSvc := svc.NewLinkSvc(linkRepo)
	linkHandler := handler.NewURLHandlerImpl(linkSvc)

	tokenSvc := svc.NewTokenSvc()
	tokenHandler := handler.NewTokenHandler(tokenSvc)
	// End of Dependency Injection

	r.Route("/url", func(r chi.Router) {
		r.Use(custommiddleware.Auth)
		// Subrouters
		r.Route("/{id}", func(r chi.Router) {
			// GET /url/123
			r.Get("/", linkHandler.Get)
			r.Put("/", linkHandler.Update)
		})
	})

	r.Get("/token", tokenHandler.Get)

	log.Println("Running on Port:", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
