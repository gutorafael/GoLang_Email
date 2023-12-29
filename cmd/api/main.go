package main

import (
	"email/internal/contract"
	"email/internal/domain/campaign"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)


func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	service := campaign.Service{}
	r.Post("/campaigns", func(w http.ResponseWriter, r *http.Request) {
		var request contract.NewCampaignDTO
		render.DecodeJSON(r.Body, &request)
		id, err := service.Create(request)
	})
	
	http.ListenAndServe(":4444", r)
}
