package handlers

import (
	"github.com/Realrubr2/ramon-coderen/golang/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(rout *chi.Mux){
	//global middleware
	rout.Use(chimiddle.StripSlashes)
	rout.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}