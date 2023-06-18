# Greenlight

Greenlight is application built while reading Let's Go Further by [Alex Edwards](https://github.com/alexedwards/)

Its a JSON API for retrieving and managing information about movies.

## Endpoints

The following endpoints are supported:

- GET /v1/healthcheck (Show application health and version information)
- GET /v1/movies (Show the details of all movies)
- POST /v1/movies (Create a new movie)
- GET /v1/movies/:id (Show the details of a specific movie)
- PATCH /v1/movies/:id (Update the details of a specific movie)
- DELETE /v1/movies/:id (Delete a specific movie)
- POST /v1/users (Register a new user)
- PUT /v1/users/activated (Activate a specific user)
- PUT /v1/users/password (Update the password for a specific user)
- POST /v1/tokens/authentication (Generate a new authentication token)
- POST /v1/tokens/password-reset (Generate a new password reset token)
- GET /debug/vars (Display application metrics)

## Sample Response

```
$ curl -H "Authorization: Bearer blah-blah-blah-blah" localhost:4000/v1/movies/1

{
    "movie": {
        "id": 1,
        "title": "The Grinch",
        "year": 2018,
        "runtime": "106 mins",
        "genres": [
            "kids",
            "family",
            "holiday"
        ],
        "version": 1
    }
}

```

## Installation

1. Clone the repo
2. `go run ./cmd/api` to start the server
3. Test the endpoints
