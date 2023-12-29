package endpoints

import (
	"net/http"
	"github.com/go-chi/render"
)

func (h *Handler) CampaignsGetAll(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, h.CampaignService.Repository.GetAll())
}
