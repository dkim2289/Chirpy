package main

import (
	"encoding/json"
	"net/http"

	"github.com/dkim2289/Chirpy/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerPolka(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil || apiKey != cfg.polkaKey {
		respondWithError(w, 401, "unauthorized")
		return
	}

	type data struct {
		UserID string `json:"user_id"`
	}
	type parameters struct {
		Event string `json:"event"`
		Data  data   `json:"data"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, "invalid JSON")
		return
	}

	if params.Event != "user.upgraded" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	userID, err := uuid.Parse(params.Data.UserID)
	if err != nil {
		respondWithError(w, 400, "Invalid user ID")
		return
	}

	err = cfg.db.UpdateUserIsChirpyRed(r.Context(), userID)
	if err != nil {
		respondWithError(w, 404, "User not found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
