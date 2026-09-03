const fs = require('fs')
const mysql = require('mysql2/promise')
const { parse } = require('csv-parse/sync')

const [songsPath, artistsPath, albumsPath] = process.argv.slice(2)

if (!songsPath || !artistsPath || !albumsPath) {
    console.error('Usage: node scripts/import-csv.js songs.csv artists.csv albums.csv')
    process.exit(1)
}

const readCsv = (filePath) => parse(fs.readFileSync(filePath), {
    relax_column_count: false,
    skip_empty_lines: true
})

const nullable = (value) => value === '' || value === '\\N' ? null : value

const pool = mysql.createPool({
    host: process.env.DB_HOST || '127.0.0.1',
    user: process.env.DB_USER,
    password: process.env.DB_PASSWORD,
    database: process.env.DB_DATABASE || 'music_db',
    port: process.env.DB_PORT || 3306
})

const importCsv = async () => {
    const artists = readCsv(artistsPath)
    const albums = readCsv(albumsPath)
    const songs = readCsv(songsPath)
    const connection = await pool.getConnection()

    try {
        await connection.beginTransaction()
        await connection.query('SET FOREIGN_KEY_CHECKS = 0')
        await connection.query('DELETE FROM songs')
        await connection.query('DELETE FROM albums')
        await connection.query('DELETE FROM artists')
        for (const row of artists) {
            await connection.query(
                'INSERT INTO artists (korean_name, id, foreign_name, debut_date, image_url) VALUES (?, ?, ?, ?, ?)',
                [row[0], row[1], row[2], nullable(row[3]), nullable(row[4])]
            )
        }

        for (const row of albums) {
            await connection.query(
                'INSERT INTO albums (title, id, korean_title, image_url, release_date, artist_id) VALUES (?, ?, ?, ?, ?, ?)',
                [row[0], row[1], nullable(row[2]), nullable(row[3]), nullable(row[5]), row[4]]
            )
        }

        for (const row of songs) {
            await connection.query(
                'INSERT INTO songs (id, title, korean_title, english_title, artist_name, song_url, image_url, release_date, disc_number, album_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)',
                [row[0], row[1], nullable(row[2]), nullable(row[3]), row[4], nullable(row[5]), nullable(row[6]), nullable(row[7]), nullable(row[8]), nullable(row[9])]
            )
        }

        await connection.query('SET FOREIGN_KEY_CHECKS = 1')
        await connection.commit()
        console.log(`Imported ${artists.length} artists, ${albums.length} albums, ${songs.length} songs.`)
    } catch (error) {
        await connection.rollback()
        throw error
    } finally {
        connection.release()
        await pool.end()
    }
}

importCsv().catch((error) => {
    console.error(error)
    process.exit(1)
})
