package handlers

import (
	"evos/be/services"
	"net/http"
)

func (h *Handler) HandleGetRoleList(responseWriter http.ResponseWriter, request *http.Request) {
	roles, err := services.GetRoleList(h.Database)
	if err != nil {
		respondWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(responseWriter, http.StatusCreated, roles)
}
