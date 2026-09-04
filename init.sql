CREATE DATABASE IF NOT EXISTS musicapp
  CHARACTER SET utf8mb4
  COLLATE utf8mb4_unicode_ci;

USE musicapp;

CREATE TABLE IF NOT EXISTS artists (
    id INT PRIMARY KEY,
    korean_name VARCHAR(255) NOT NULL,
    foreign_name VARCHAR(255) NOT NULL,
    debut_date DATE NULL,
    image_url VARCHAR(512) NULL
)CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS albums (
    id INT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    artist_id INT NOT NULL,
    korean_title VARCHAR(255) NULL,
    image_url VARCHAR(512) NULL,
    release_date DATE NULL,
    INDEX idx_albums_artist_id (artist_id),
    CONSTRAINT fk_albums_artist FOREIGN KEY (artist_id) REFERENCES artists(id)
)CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS songs (
    id INT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    korean_title VARCHAR(255) NULL,
    english_title VARCHAR(255) NULL,
    artist_id INT NOT NULL,
    song_url VARCHAR(1024) NULL,
    image_url VARCHAR(1024) NULL,
    release_date DATE NULL,
    disc_number INT NULL,
    album_id INT NULL,
    INDEX idx_songs_artist_id (artist_id),
    INDEX idx_songs_album_id (album_id),
    CONSTRAINT fk_songs_artist FOREIGN KEY (artist_id) REFERENCES artists(id),
    CONSTRAINT fk_songs_album FOREIGN KEY (album_id) REFERENCES albums(id)
)CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
