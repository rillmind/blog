# Blog

Personal project of a blog that I post about my life and my opinions

## What it does

A Go + Gin REST API (port `2306`) that stores blog posts in PostgreSQL through GORM, and a React +
Vite front-end (port `2207`). Current routes:

| Method | Route    | Action          |
| ------ | -------- | --------------- |
| GET    | `/posts` | list all posts  |
| POST   | `/posts` | create a post (`{"title": "...", "content": "..."}`) |

User and product modules exist in `back-end/` but are not wired into `cmd/main.go` yet.

## Why it exists

I wanted my own place to write about my life and opinions, built with the stack I was studying:
Go with Gin on the back-end and React on the front-end.

## Stack

- Gin gonic for back-end
- React for front-end

## Run

On linux based OS and docker installed run:

```sh
make run
```

## Installation

Requirements: Go 1.24+, [Bun](https://bun.sh), Podman + podman-compose (the makefile calls
`podman`, not `docker`) and `make`.

```sh
git clone https://github.com/rillmind/blog.git
cd blog
podman-compose up -d db          # PostgreSQL 12 on localhost:5432 (postgres / 1234)
MIGRATION=true make b            # back-end on :2306; MIGRATION=true creates the posts table
make f                           # in another terminal: installs dependencies and runs Vite on :2207
```

`make run` builds both images and starts the whole compose stack, but in a clean clone the `back`
container exits right away: inside the container the default database host is `127.0.0.1`, and
the compose file sets neither `DATABASE_URL` nor `MIGRATION`. The `db` and `front` containers do
start. Until the compose file passes those variables, run the back-end on the host as shown above.

If Podman fails with "short-name resolution", pull the images by their full names first:
`podman pull docker.io/library/postgres:12 docker.io/library/golang:1.24.3 docker.io/oven/bun`.

## How to test

There are no automated tests yet. To check that everything builds:

```sh
(cd back-end && go vet ./...)
(cd front-end && bun install && bun run build && bun run lint)
```

With the database and the back-end running:

```sh
curl -X POST localhost:2306/posts -H 'content-type: application/json' \
  -d '{"title":"Hello","content":"first post"}'   # 201 with the created post
curl localhost:2306/posts                         # list containing the post
```

Then open `http://localhost:2207` and check that the page loads.
