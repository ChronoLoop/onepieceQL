# OnePieceQL
![One_Piece_Anime_Logo](https://github.com/ChronoLoop/onepieceQL/assets/44091034/43a9d57d-fbc3-4bff-8516-8f35941e7e44)

## Introduction

OnePieceQL is a GraphQL API that provides extensive data about the One Piece universe, a popular manga and anime series.

You can view the live [documentation](https://onepieceql.up.railway.app)

## Getting Started

Ensure [PNPM](https://pnpm.io/installation) and [Docker](https://docs.docker.com/engine/install) are installed.

Create .env file and add environment variables ([see below](#environment-variables))

Run docker compose

```bash
docker compose up
```

### Server

```bash
cd ./server
```

Download dependencies

```bash
go mod download
```

Run server

```bash
go run ./cmd/server/main.go
```

### Client

```bash
cd ./client
```

Download dependencies

```bash
pnpm i --frozen-lockfile
```

Run client

```bash
pnpm run dev
```

## Environment Variables
```bash
POSTGRES_USER=?
POSTGRES_PASSWORD=?
POSTGRES_DB=?
POSTGRES_PORT=5432
POSTGRES_HOST=localhost
SERVER_ENV=development
PORT=5000

AWS_REGION=?
AWS_ACCESS_KEY_ID=?
AWS_SECRET_ACCESS_KEY=?
AWS_S3_BUCKET_NAME=?
```
