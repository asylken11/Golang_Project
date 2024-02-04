package main

import (
	"encoding/json"
	"net/http"
)

// NaVi is an instance of Team for Natus Vincere
var NaVi = Team{
	Name:    "Natus Vincere",
	Country: "Ukraine",
	Players: []Player{
		{"w0nderful", "AWPer", "Ukraine"},
		{"jL", "Rifler", "Lithuania"},
		{"Aleksib", "In-Game Leader", "Finland"},
		{"b1t", "Rifler", "Ukraine"},
		{"iM", "Rifler", "Romania"},
	},
}

// Team represents a CS:GO team
type Team struct {
	Name    string   `json:"name"`
	Country string   `json:"country"`
	Players []Player `json:"players"`
}

func getTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(NaVi)
}
