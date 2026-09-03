const express = require('express')
const mysql = require('mysql2/promise')
const router = express.Router()
require('dotenv').config()

const pool = mysql.createPool({
    host: process.env.DB_HOST,
    user: process.env.DB_USER,
    password: process.env.DB_PASSWORD,
    database: process.env.DB_DATABASE || 'music_db',
    waitForConnections: true,
    connectionLimit: 10,
    queueLimit: 0
})

router.use('/img/artist', (req, res) => {
    res.redirect(`https://storage.kimrasng.kr/music_server/img${req.path}`)
})
router.use('/img/album', (req, res) => {
    res.redirect(`https://storage.kimrasng.kr/music_server/img${req.path}`)
})
router.use('/img/song', (req, res) => {
    res.redirect(`https://storage.kimrasng.kr/music_server/img${req.path}`)
})
router.use('/music', (req, res) => {
    res.redirect(`https://storage.kimrasng.kr/music_server/songs${req.path}`)
})

const getRequestValue = (req, name) => {
    const bodyValue = req.body && req.body[name]
    return bodyValue || req.get(name) || req.get(`x-${name}`)
}

const getSortOrder = (req) => ({
    sort: getRequestValue(req, 'sort'),
    order: getRequestValue(req, 'order') || 'asc'
})

// music-api/songs
router.get('/songs', async (req, res) => {
    try {
        const { sort, order } = getSortOrder(req)
        const sortColumns = { title: 'title', artist: 'artist_name', date: 'release_date' }
        const orderBy = sortColumns[sort]

        if (sort && !orderBy) return res.status(400).send('Invalid sort parameter')
        if (!['asc', 'desc'].includes(order)) return res.status(400).send('Invalid order parameter')

        const query = orderBy
            ? `SELECT * FROM songs ORDER BY ${orderBy} ${order.toUpperCase()}`
            : 'SELECT * FROM songs'
        const [rows] = await pool.query(query)
        res.json({ songs: rows })
    } catch (error) {
        console.error(error)
        res.status(500).send('Error fetching songs')
    }
})

router.get('/songs/:id', async (req, res) => {
    try {
        const [rows] = await pool.query('SELECT * FROM songs WHERE id = ?', [req.params.id])
        if (!rows.length) return res.status(404).send('Song not found')
        res.json({ song: rows[0] })
    } catch (error) {
        console.error(error)
        res.status(500).send('Error fetching song')
    }
})

// music-api/artists
router.get('/artists', async (req, res) => {
    try {
        const { sort, order } = getSortOrder(req)
        const sortColumns = { name: 'korean_name', date: 'debut_date', foreign: 'foreign_name' }
        const orderBy = sortColumns[sort]

        if (sort && !orderBy) return res.status(400).send('Invalid sort parameter')
        if (!['asc', 'desc'].includes(order)) return res.status(400).send('Invalid order parameter')

        const query = orderBy
            ? `SELECT * FROM artists ORDER BY ${orderBy} ${order.toUpperCase()}`
            : 'SELECT * FROM artists'
        const [rows] = await pool.query(query)
        res.json({ artists: rows })
    } catch (error) {
        console.error(error)
        res.status(500).send('Error fetching artists')
    }
})

router.get('/artists/:id', async (req, res) => {
    try {
        const [rows] = await pool.query('SELECT * FROM artists WHERE id = ?', [req.params.id])
        if (!rows.length) return res.status(404).send('Artist not found')
        res.json({ artist: rows[0] })
    } catch (error) {
        console.error(error)
        res.status(500).send('Error fetching artist')
    }
})

router.get('/albums', async (req, res) => {
    try {
        const { sort, order } = getSortOrder(req)
        const sortColumns = { title: 'title', artist: 'artist_id', date: 'release_date' }
        const orderBy = sortColumns[sort]

        if (sort && !orderBy) return res.status(400).send('Invalid sort parameter')
        if (!['asc', 'desc'].includes(order)) return res.status(400).send('Invalid order parameter')

        const query = orderBy
            ? `SELECT * FROM albums ORDER BY ${orderBy} ${order.toUpperCase()}`
            : 'SELECT * FROM albums'
        const [rows] = await pool.query(query)
        res.json({ albums: rows })
    } catch (error) {
        console.error(error)
        res.status(500).send('Error fetching albums')
    }
})

router.get('/albums/:id', async (req, res) => {
    try {
        const [rows] = await pool.query('SELECT * FROM albums WHERE id = ?', [req.params.id])
        if (!rows.length) return res.status(404).send('Album not found')
        res.json({ album: rows[0] })
    } catch (error) {
        console.error(error)
        res.status(500).send('Error fetching album')
    }
})

module.exports = router
