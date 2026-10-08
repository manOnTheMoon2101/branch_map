package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Branch struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`

	City      string  `json:"city"`
	Province  string  `json:"province"`
	Phone     string  `json:"phone"`
	Hours     string  `json:"hours"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

var branches = []Branch{
	{1, "Capitec Bank Wellington", "branch", "Wellington", "Western Cape", "+27 860 102 043", "Mon–Fri 08:00–16:30 | Sat 08:00–13:00", -33.639735, 19.008627},
	{2, "Capitec Bank Paarl Main Road", "branch", "Paarl", "Western Cape", "+27 860 102 043", "Mon–Fri 08:00–16:30 | Sat 08:00–13:00", -33.731264, 18.963111},
	{3, "Capitec Bank Paarl Mall", "branch", "Paarl", "Western Cape", "+27 21 941 1377", "Mon–Fri 08:00–16:30 | Sat 08:00–13:00", -33.765054, 18.968120},
	{4, "Capitec Bank V&A Waterfront", "branch", "Cape Town", "Western Cape", "+27 860 102 043", "Mon–Fri 08:00–16:30 | Sat 08:00–13:00", -33.904462, 18.419247},
	{5, "Capitec Bank Head Office", "head", "Stellenbosch", "Western Cape", "+27 218 09 500", "Mon–Fri 09:00–16:00 | Sat 08:00–13:00", -33.964493, 18.832710},
	{6, "Capitec Bank Stellenbosch Eikestad Mall", "branch", "Stellenbosch", "Western Cape", "+27 860 102 043", "Mon–Fri 09:00–18:00 | Sat 08:00–13:00", -33.934996, 18.860101},
	{7, "Capitec Bank Malmesbury Voortrekker", "branch", "Malmesbury", "Western Cape", "+27 860 102 043", "Mon–Fri 09:00–18:00 | Sat 08:00–13:00", -33.462407, 18.729089},
	{8, "Capitec Bank Mossel Bay Langeberg", "branch", "Mossel Bay", "Western Cape", "+27 860 102 043", "Mon–Fri 09:00–18:00 | Sat 08:00–13:00", -34.148369, 22.103822},
	{9, "Capitec Bank Worcester Mountain Mill Centre", "branch", "Worcester", "Western Cape", "+27 860 102 043", "Mon–Fri 09:00–18:00 | Sat 08:00–13:00", -33.632267, 19.434541},
	{10, "Capitec Bank Somerset Mall", "branch", "Somerset West", "Western Cape", "+27 860 102 043", "Mon–Fri 09:00–18:00 | Sat 08:00–13:00", -34.080431, 18.822535},
}

func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next(w, r)
	}
}

func branchesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(branches); err != nil {
		log.Printf("error encoding branches: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/branches", withCORS(branchesHandler))
	addr := ":8080"
	log.Printf("Branch API listening on http://localhost%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
