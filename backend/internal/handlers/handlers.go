package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

// Health returns 200 when the service is up (no DB check).
func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

// Ready returns 200 if the database is reachable, 503 otherwise.
func Ready(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if err := db.Ping(); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte("database unreachable"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ready"))
	}
}

func URLs(db *sql.DB) http.HandlerFunc {
	type RequestBody struct {
		URL string `json:"url"`
	}
	type ResponseBody struct {
		ID  int64  `json:"id"`
		URL string `json:"url"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		defer r.Body.Close()

		var body RequestBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("request body must be valid JSON"))
			return
		}
		if body.URL == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("url is required"))
			return
		}

		log.Printf("inserting url: %s", body.URL)
		result, err := db.Exec("INSERT INTO urls (url) VALUES (?)", body.URL)
		if err != nil {
			log.Printf("failed to insert url: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to insert url"))
			return
		}

		id, err := result.LastInsertId()
		if err != nil {
			log.Printf("failed to get inserted id: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to insert url"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		resp := ResponseBody{
			ID:  id,
			URL: body.URL,
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Printf("failed to encode response: %v", err)
		}
	}
}
