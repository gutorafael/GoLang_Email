package endpoints

import (
	"net/http"
)

func (h *Handler) CampaignsGetAll(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	campaigns, err := h.CampaignService.Repository.GetAll()
	return campaigns, http.StatusOK, err
}
