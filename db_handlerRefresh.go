package main

import (
	"net/http"
	"time"

	"github.com/dkim2289/Chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	tokenString, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, 401, "Unauthorized")
		return
	}

	refreshToken, err := cfg.db.GetUserFromRefreshToken(r.Context(), tokenString)
	if err != nil {
		respondWithError(w, 401, "Unauthorized")
		return
	}

	if refreshToken.RevokedAt.Valid || time.Now().UTC().After(refreshToken.ExpiresAt) {
		respondWithError(w, 401, "Unauthorized")
		return
	}

	accessToken, err := auth.MakeJWT(refreshToken.UserID, cfg.jwtSecret, time.Hour)
	if err != nil {
		respondWithError(w, 500, "Could not create access token")
		return
	}

	type response struct {
		Token string `json:"token"`
	}
	respondWithJSON(w, 200, response{Token: accessToken})
}

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	tokenString, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, 401, "Unauthorized")
		return
	}

	err = cfg.db.RevokeRefreshToken(r.Context(), tokenString)
	if err != nil {
		respondWithError(w, 500, "COuld not revoke token")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
