package models

type Artist struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	KoreanName  string  `json:"korean_name"`
	ForeignName string  `json:"foreign_name"`
	DebutDate   *string `json:"debut_date"`
	ImageURL    *string `json:"image_url"`
}

type Album struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	ArtistID    int     `json:"artist_id"`
	ArtistName  string  `json:"artist_name"`
	KoreanTitle *string `json:"korean_title"`
	ImageURL    *string `json:"image_url"`
	ReleaseDate *string `json:"release_date"`
}

type Song struct {
	ID           int     `json:"id"`
	Title        string  `json:"title"`
	KoreanTitle  *string `json:"korean_title"`
	EnglishTitle *string `json:"english_title"`
	ArtistID     int     `json:"artist_id"`
	ArtistName   string  `json:"artist_name"`
	SongURL      *string `json:"song_url"`
	ImageURL     *string `json:"image_url"`
	ReleaseDate  *string `json:"release_date"`
	DiscNumber   *int    `json:"disc_number"`
	AlbumID      *int    `json:"album_id"`
	AlbumTitle   *string `json:"album_title"`
}
