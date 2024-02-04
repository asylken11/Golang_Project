package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Player represents a CS:GO player
type Player struct {
	Name   string `json:"name"`
	Role   string `json:"role"`
	Nation string `json:"nation"`
}

func getPlayerByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	playerName := params["name"]

	for _, player := range NaVi.Players {
		if player.Name == playerName {
			json.NewEncoder(w).Encode(player)
			return
		}
	}

	http.Error(w, "Player not found", http.StatusNotFound)
}
