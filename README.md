# Prerequisites
Before going any further, make sure that you have Postgresql running. Take a note of its:
- Hostname
- Port
- username
- password
- database name

# Database Migration
We're using [Tern](https://github.com/jackc/tern) to perform database migration. Go there and install Tern. Once we're done, head over to `/db/migrations/tern.conf` and fill it with database config.

Once that is done change terminal directory to `/db/migrations` and type `tern migrate` to start performing database migration.

If successful, our database should be populated by tables and dummy datas.

# Environment variable
Make a copy of `.env.example`, rename it to `.env` and populate it with database configuration and our own app configuration, along with Secret key used for generating JWT Token.

# Running the app
Type `go run main.go` to start Gin.

# Authenticating
Peform a **POST** request to `/auth/` with username and password as JSON body. By default, it comes with dummy account as part of database migration. We can use `admin` both as its username and password.

If successful, JWT Token will stored as cookies with expiry date of one hour from token generation.

# Available API
`GET /health-check/` -> used to ping the server. Token is not required to perform this check.

## Authentication
`POST /auth/` -> used to authenticate and generate JWT Token. Requires JSON with admin and password field.

`POST /auth/verify-token/` -> Check whether token is still valid or not, requires token inside cookie.

`POST /auth/logout` -> Log users out by removing their JWT Token from cookies.

## Movies
`POST /api/v1/movie/` -> insert new movie. Data must be sent as JSON that matches the data model located at `/model/movie-model.go`.

`GET /api/v1/movie/:id/` -> fetch movie data based off its ID. The ID is sent as url parameter.

`DELETE /api/v1/movie/:id/` -> delete movie data based off its ID. The ID is sent as url parameter.

`GET /api/v1/movie/byPublisher/:id/` -> fetch one or more movies data based off its publisher ID. The ID is sent as url parameter.

`GET /api/v1/movies/` -> fetch all movies.

## Publishers
`GET /api/v1/publishers/name/` -> fetch all publishers. This endpoint only fetches publisher's name and its ID.

## Genre
`GET /api/v1/genres/` -> fetch all genres.
