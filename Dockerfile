FROM node:22-alpine AS builder-web
WORKDIR /app
COPY web/package.json web/package-lock.json ./
RUN npm ci
COPY web/ ./
RUN npm run build

FROM golang:1.24-alpine AS builder-go
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY --from=builder-web /app/dist/ web/dist/
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /moyoki ./cmd/server

FROM alpine

RUN apk add --no-cache tzdata
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
RUN mkdir /config && chown appuser:appgroup /config
USER appuser

ENV PANEL_LISTEN_ADDR=:3000
ENV MOCK_LISTEN_ADDR=:3001
ENV DB_DSN=/config/moyoki.db
ENV TZ=Etc/UTC

VOLUME /config

COPY --from=builder-go /moyoki /usr/local/bin/moyoki

EXPOSE 3000 3001

ENTRYPOINT ["moyoki"]