FROM node:20.11.1 AS clientBuilder
WORKDIR /app

COPY ./client/. /app
RUN npm i -g pnpm 
RUN pnpm i --frozen-lockfile
RUN pnpm run build



FROM golang:1.22.1-alpine AS serverBuilder
WORKDIR /app

COPY ./server/. /app
RUN go mod download
RUN GOOS=linux go build -ldflags="-s -w" -o app ./cmd/server/main.go



FROM alpine:latest
WORKDIR /app

RUN mkdir ./client
RUN mkdir ./client/dist/
RUN mkdir ./server

COPY --from=serverBuilder /app/app ./server
COPY --from=serverBuilder /app/internal/server/db/migrations ./server/internal/server/db/migrations/
COPY --from=serverBuilder /app/data ./server/data/
COPY --from=clientBuilder /app/dist ./client/dist/

EXPOSE 5000

WORKDIR /app/server
CMD ["./app"]
