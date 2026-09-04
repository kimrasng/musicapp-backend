package handlers

import (
	"database/sql"
	"net/http"

	"musicapp-backend/app/models"
)

func (h *Handler) Songs(w http.ResponseWriter, r *http.Request) {
	options, err := parseListQuery(r, "id", "asc", map[string]string{
		"id": "songs.id", "title": "songs.title", "release_date": "songs.release_date",
	})
	if err != nil {
		queryError(w, err)
		return
	}

	query := `
		SELECT songs.id, songs.title, songs.korean_title, songs.english_title, songs.artist_id, artists.korean_name,
		       songs.song_url, songs.image_url, songs.release_date, songs.disc_number, songs.album_id, albums.title
		FROM songs JOIN artists ON artists.id = songs.artist_id LEFT JOIN albums ON albums.id = songs.album_id
		`
	args := []any{}
	if options.ID != nil {
		query += " WHERE songs.id = ?"
		args = append(args, *options.ID)
	}
	query += " ORDER BY " + options.SortValue + " " + options.Order
	rows, err := h.db.QueryContext(r.Context(), query, args...)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not load songs")
		return
	}
	defer rows.Close()

	songs := []models.Song{}
	for rows.Next() {
		var song models.Song
		var koreanTitle, englishTitle, songURL, imageURL, releaseDate, albumTitle sql.NullString
		var discNumber, albumID sql.NullInt64
		if err := rows.Scan(&song.ID, &song.Title, &koreanTitle, &englishTitle, &song.ArtistID, &song.ArtistName, &songURL, &imageURL, &releaseDate, &discNumber, &albumID, &albumTitle); err != nil {
			writeError(w, http.StatusInternalServerError, "could not read songs")
			return
		}
		song.KoreanTitle = nullStringValue(koreanTitle)
		song.EnglishTitle = nullStringValue(englishTitle)
		song.SongURL = nullStringValue(songURL)
		song.ImageURL = nullStringValue(imageURL)
		song.ReleaseDate = nullStringValue(releaseDate)
		song.DiscNumber = nullIntValue(discNumber)
		song.AlbumID = nullIntValue(albumID)
		song.AlbumTitle = nullStringValue(albumTitle)
		songs = append(songs, song)
	}
	if err := rows.Err(); err != nil {
		writeError(w, http.StatusInternalServerError, "could not read songs")
		return
	}
	writeJSON(w, http.StatusOK, songs)
}
