package endpoints

import (
	"email/internal/contract"
	internalerrors "email/internal/internalErrors"
	"errors"
	"net/http"
	"github.com/go-chi/render"
)

func (h *Handler) CampaignsPost(w http.ResponseWriter, r *http.Request) {
	var request contract.NewCampaignDTO
	render.DecodeJSON(r.Body, &request)
	id, err := h.CampaignService.Create(request)

	if err != nil {

		if errors.Is(err, internalerrors.ErrInternal) {
			render.Status(r, http.StatusInternalServerError)

		} else {
			render.Status(r, http.StatusBadRequest)

		}
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]string{"campaign_id": id})
}
