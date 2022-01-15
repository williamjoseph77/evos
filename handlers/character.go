package handlers

import (
	"encoding/json"
	"evos/be/objects"
	"evos/be/services"
	"fmt"
	"net/http"
)

func (h *Handler) HandleCreateCharacterNonSecure(responseWriter http.ResponseWriter, request *http.Request) {
	var requestPayload objects.CreateCharacterNonSecureRequest
	var responsePayload objects.CreateCharacterResponse

	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&requestPayload); err != nil {
		respondWithError(responseWriter, http.StatusBadRequest, "Invalid request payload")
		return
	}
	fmt.Println("request payload")
	fmt.Println(requestPayload)
	defer request.Body.Close()

	character, err := services.CreateCharacterNonSecure(h.Database, requestPayload)
	if err != nil {
		respondWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	responsePayload.GUID = character.GUID
	fmt.Println("responsenya harusnya")
	fmt.Println(responsePayload)
	respondWithJSON(responseWriter, http.StatusCreated, responsePayload)
}

func (h *Handler) HandleCreateCharacterSecure(responseWriter http.ResponseWriter, request *http.Request) {
	var requestPayload objects.CreateCharacterSecureRequest
	var responsePayload objects.CreateCharacterResponse

	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&requestPayload); err != nil {
		respondWithError(responseWriter, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer request.Body.Close()

	character, err := services.CreateCharacterSecure(h.Database, requestPayload)
	if err != nil {
		respondWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	responsePayload.GUID = character.GUID

	respondWithJSON(responseWriter, http.StatusCreated, responsePayload)
}

func (h *Handler) HandleGetCharacterList(responseWriter http.ResponseWriter, request *http.Request) {
	characters, err := services.GetCharacterList(h.Database)
	if err != nil {
		respondWithError(responseWriter, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(responseWriter, http.StatusCreated, characters)
}
