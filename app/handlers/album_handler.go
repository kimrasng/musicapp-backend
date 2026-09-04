package handlers

import (
	"database/sql"
	"net/http"

	"musicapp-backend/app/models"
)

func (h *Handler) Albums(w http.ResponseWriter, r *http.Request) {
	options, err := parseListQuery(r, "release_date", "desc", map[string]string{
		"id": "albums.id", "title": "albums.title", "release_date": "albums.release_date",
	})
	if err != nil {
		queryError(w, err)
		return
	}

	query := `
		SELECT albums.id, albums.title, albums.artist_id, artists.korean_name, albums.korean_title, albums.image_url, albums.release_date
		FROM albums JOIN artists ON artists.id = albums.artist_id`
	args := []any{}
	if options.ID != nil {
		query += " WHERE albums.id = ?"
		args = append(args, *options.ID)
	}
	query += " ORDER BY " + options.SortValue + " " + options.Order
	rows, err := h.db.QueryContext(r.Context(), query, args...)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not load albums")
		return
	}
	defer rows.Close()

	albums := []models.Album{}
	for rows.Next() {
		var album models.Album
		var koreanTitle, imageURL, releaseDate sql.NullString
		if err := rows.Scan(&album.ID, &album.Title, &album.ArtistID, &album.ArtistName, &koreanTitle, &imageURL, &releaseDate); err != nil {
			writeError(w, http.StatusInternalServerError, "could not read albums")
			return
		}
		album.KoreanTitle = nullStringValue(koreanTitle)
		album.ImageURL = nullStringValue(imageURL)
		album.ReleaseDate = nullStringValue(releaseDate)
		albums = append(albums, album)
	}
	if err := rows.Err(); err != nil {
		writeError(w, http.StatusInternalServerError, "could not read albums")
		return
	}
	writeJSON(w, http.StatusOK, albums)
}
