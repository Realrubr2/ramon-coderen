package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/realrubr/golang/internal/middleware"
)

func handler(rout *chi.Mux){
	//global middleware
	rout.Use(chimiddle.StripSlashes)
	rout.Route("/account", func(router chi.Router) {
		router.use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}