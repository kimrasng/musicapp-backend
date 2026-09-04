package app

import (
	"database/sql"
	"net/http"

	"musicapp-backend/app/handlers"
)

func ListenAndServe(db *sql.DB, address string) error {
	return http.ListenAndServe(address, NewHandler(db))
}

func NewHandler(db *sql.DB) http.Handler {
	api := handlers.New(db)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", api.Health)
	mux.HandleFunc("GET /api/artists", api.Artists)
	mux.HandleFunc("GET /api/albums", api.Albums)
	mux.HandleFunc("GET /api/songs", api.Songs)
	mux.HandleFunc("GET /openapi.json", swaggerJSON)
	mux.HandleFunc("GET /swagger", swaggerUI)
	return cors(mux)
}
