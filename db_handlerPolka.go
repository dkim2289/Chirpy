package main

import "net/http"

type polkaRequest struct {
	event string `json:"event"`
	data  struct {
		user_id string `json:"user_id"`
	} `json:"data"`
}

func handlerPolka(w http.ResponseWriter, pr polkaRequest) {
	if pr.event != "user.upgraded" {
		respondWithError(w, 204, "user not upgraded")
		return
	}

	err := db.UpgradeUser(pr.data.user_id)
	if err != nil {
		respondWithError(w, 500, err.Error())
		return
	}

	respondWithJSON()

}
