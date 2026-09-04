package handlers

import (
	"database/sql"
	"net/http"

	"musicapp-backend/app/models"
)

func (h *Handler) Artists(w http.ResponseWriter, r *http.Request) {
	rows, err := h.db.QueryContext(r.Context(), "SELECT id, korean_name, foreign_name, debut_date, image_url FROM artists ORDER BY korean_name")
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not load artists")
		return
	}
	defer rows.Close()

	artists := []models.Artist{}
	for rows.Next() {
		var artist models.Artist
		var debutDate, imageURL sql.NullString
		if err := rows.Scan(&artist.ID, &artist.KoreanName, &artist.ForeignName, &debutDate, &imageURL); err != nil {
			writeError(w, http.StatusInternalServerError, "could not read artists")
			return
		}
		artist.Name = artist.KoreanName
		artist.DebutDate = nullStringValue(debutDate)
		artist.ImageURL = nullStringValue(imageURL)
		artists = append(artists, artist)
	}
	if err := rows.Err(); err != nil {
		writeError(w, http.StatusInternalServerError, "could not read artists")
		return
	}
	writeJSON(w, http.StatusOK, artists)
}
