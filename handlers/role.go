package handlers

import (
	"net/http"

	"github.com/williamjoseph77/evos/services"
)

func (h *Handler) HandleGetRoleList(responseWriter http.ResponseWriter, request *http.Request) {
	roles, err := services.GetRoleList(h.Database)
	if err != nil {
		respondWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(responseWriter, http.StatusCreated, roles)
}
