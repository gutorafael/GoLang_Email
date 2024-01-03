package main

import (
	"email/internal/domain/campaign"
	"email/internal/endpoints"
	"email/internal/infrastructure/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaignService := campaign.Service{
		Repository: &database.CampaignRepo{},
	}
	handler := endpoints.Handler{
		CampaignService: campaignService,
	}

	r.Post("/campaigns", endpoints.HandlerError(handler.CampaignsPost))
	r.Get("/campaigns", endpoints.HandlerError(handler.CampaignsGetAll))
	http.ListenAndServe(":4444", r)
}
