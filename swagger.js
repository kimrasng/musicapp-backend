const swaggerJSDoc = require('swagger-jsdoc')

const options = {
    definition: {
        openapi: '3.0.3',
        info: {
            title: 'HomeMusic API',
            version: '1.0.0',
            description: 'Music, artist, album, media redirect, and database test APIs.'
        },
        servers: [
            { url: '/api', description: 'API server' }
        ],
        tags: [
            { name: 'Songs', description: 'Song list and details' },
            { name: 'Artists', description: 'Artist list and details' },
            { name: 'Albums', description: 'Album list and details' },
            { name: 'Media', description: 'Image and music redirects' },
            { name: 'Test', description: 'Database connectivity test' }
        ],
        components: {
            parameters: {
                Sort: {
                    name: 'sort',
                    in: 'header',
                    required: false,
                    schema: { type: 'string' },
                    description: 'Sort key. Songs/albums: title, artist, date. Artists: name, foreign, date.'
                },
                Order: {
                    name: 'order',
                    in: 'header',
                    required: false,
                    schema: { type: 'string', enum: ['asc', 'desc'], default: 'asc' }
                }
            },
            schemas: {
                Song: { type: 'object', additionalProperties: true },
                Artist: { type: 'object', additionalProperties: true },
                Album: { type: 'object', additionalProperties: true }
            }
        },
        paths: {
            '/music-server/songs': {
                get: {
                    tags: ['Songs'],
                    summary: 'Get all songs',
                    parameters: [
                        { $ref: '#/components/parameters/Sort' },
                        { $ref: '#/components/parameters/Order' }
                    ],
                    responses: { 200: { description: 'Song list' }, 400: { description: 'Invalid sort or order' } }
                }
            },
            '/music-server/songs/{id}': {
                get: {
                    tags: ['Songs'],
                    summary: 'Get a song by ID',
                    parameters: [{ $ref: '#/components/parameters/Id' }],
                    responses: { 200: { description: 'Song detail' }, 404: { description: 'Song not found' } }
                }
            },
            '/music-server/artists': {
                get: {
                    tags: ['Artists'],
                    summary: 'Get all artists',
                    parameters: [
                        { $ref: '#/components/parameters/Sort' },
                        { $ref: '#/components/parameters/Order' }
                    ],
                    responses: { 200: { description: 'Artist list' }, 400: { description: 'Invalid sort or order' } }
                }
            },
            '/music-server/artists/{id}': {
                get: {
                    tags: ['Artists'],
                    summary: 'Get an artist by ID',
                    parameters: [{ $ref: '#/components/parameters/Id' }],
                    responses: { 200: { description: 'Artist detail' }, 404: { description: 'Artist not found' } }
                }
            },
            '/music-server/albums': {
                get: {
                    tags: ['Albums'],
                    summary: 'Get all albums',
                    parameters: [
                        { $ref: '#/components/parameters/Sort' },
                        { $ref: '#/components/parameters/Order' }
                    ],
                    responses: { 200: { description: 'Album list' }, 400: { description: 'Invalid sort or order' } }
                }
            },
            '/music-server/albums/{id}': {
                get: {
                    tags: ['Albums'],
                    summary: 'Get an album by ID',
                    parameters: [{ $ref: '#/components/parameters/Id' }],
                    responses: { 200: { description: 'Album detail' }, 404: { description: 'Album not found' } }
                }
            },
            '/music-server/img/artist/{path}': { get: { tags: ['Media'], summary: 'Redirect to artist image', parameters: [{ $ref: '#/components/parameters/Path' }], responses: { 302: { description: 'Image redirect' } } } },
            '/music-server/img/album/{path}': { get: { tags: ['Media'], summary: 'Redirect to album image', parameters: [{ $ref: '#/components/parameters/Path' }], responses: { 302: { description: 'Image redirect' } } } },
            '/music-server/img/song/{path}': { get: { tags: ['Media'], summary: 'Redirect to song image', parameters: [{ $ref: '#/components/parameters/Path' }], responses: { 302: { description: 'Image redirect' } } } },
            '/music-server/music/{path}': { get: { tags: ['Media'], summary: 'Redirect to music file', parameters: [{ $ref: '#/components/parameters/Path' }], responses: { 302: { description: 'Music redirect' } } } },
            '/test/db': { get: { tags: ['Test'], summary: 'Test database connection', responses: { 200: { description: 'Database connection succeeded' }, 500: { description: 'Database connection failed' } } } }
        }
    },
    apis: []
}

options.definition.components.parameters.Id = {
    name: 'id',
    in: 'path',
    required: true,
    schema: { type: 'integer', example: 1 }
}
options.definition.components.parameters.Path = {
    name: 'path',
    in: 'path',
    required: true,
    schema: { type: 'string', example: 'cover.jpg' }
}

module.exports = swaggerJSDoc(options)
