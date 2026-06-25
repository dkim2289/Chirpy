# Chirpy

A RESTful HTTP server written in Go, built as a Twitter-like microblogging API. Chirpy supports user authentication, chirp (post) management, and a membership system called Chirpy Red.

## Features

- User registration and login with password hashing (Argon2)
- JWT-based authentication with refresh token support
- Create, read, and delete chirps (max 140 characters, profanity filtered)
- Chirpy Red membership via webhook integration (Polka)
- PostgreSQL database with migrations (Goose) and type-safe queries (SQLC)
- Admin endpoints for metrics and database reset

## Tech Stack

- **Language**: Go
- **Database**: PostgreSQL
- **Migrations**: Goose
- **Query generation**: SQLC
- **Auth**: JWT (golang-jwt), Argon2id (alexedwards/argon2id)

## Requirements

- Go 1.22+
- PostgreSQL
- Goose (`go install github.com/pressly/goose/v3/cmd/goose@latest`)

## Installation

1. Clone the repository:
```sh
   git clone https://github.com/dkim2289/Chirpy.git
   cd Chirpy
```

2. Install dependencies:
```sh
   go mod tidy
```

3. Create a `.env` file in the project root:
```sh
    DB_URL=postgres://your_user:@localhost:5432/chirpy?sslmode=disable
    PLATFORM=dev
    JWT_SECRET=your_jwt_secret
    POLKA_KEY=your_polka_key
```

4. Run database migrations from `sql/schema`:
```sh
   cd sql/schema
   goose postgres "postgres://your_user:@localhost:5432/chirpy" up
   cd ../..
```

5. Build and run:
```sh
   go build -o Chirpy && ./Chirpy
```

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| POST | `/api/users` | Register a new user |
| PUT | `/api/users` | Update email and password (auth required) |
| POST | `/api/login` | Login and receive access + refresh tokens |
| POST | `/api/refresh` | Get a new access token via refresh token |
| POST | `/api/revoke` | Revoke a refresh token |
| POST | `/api/chirps` | Create a chirp (auth required) |
| GET | `/api/chirps` | Get all chirps (supports `author_id` and `sort` query params) |
| GET | `/api/chirps/{chirpID}` | Get a single chirp |
| DELETE | `/api/chirps/{chirpID}` | Delete a chirp (auth required, must be author) |
| GET | `/api/healthz` | Health check |
| POST | `/api/polka/webhooks` | Polka webhook for Chirpy Red upgrades |
| GET | `/admin/metrics` | View request count |
| POST | `/admin/reset` | Reset database (dev only) |
