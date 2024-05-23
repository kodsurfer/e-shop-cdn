# Base Stage
FROM golang:1.22-alpine AS base
LABEL maintainer="Kartashov Egor kartashov_egor96@mail.ru"
# if use private libs
#ARG GITHUB_TOKEN
#RUN apk update && apk add ca-certificates git openssh
#RUN git config --global url."https://${GITHUB_TOKEN}@github.com/".insteadOf "https://github.com/"
WORKDIR /srv
COPY go.mod go.sum ./
RUN go mod download && mkdir -p dist

# Development Stage
FROM base as dev
WORKDIR /srv/
COPY . .
RUN go install -mod=mod github.com/cosmtrek/air
CMD ["air", "-c", ".air-unix.toml", "-d"]

# # Test Stage
# FROM base as test
# ENTRYPOINT make test

# Build Production Stage
FROM base as builder
WORKDIR /srv
COPY . .
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN $GOPATH/bin/swag init -g cmd/main.go --output docs
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o dist/app cmd/main.go

# Production Stage
FROM cgr.dev/chainguard/busybox:latest-glibc as production
WORKDIR /srv/
COPY --from=builder /srv/docs/* ./docs
COPY --from=builder /srv/dist/app ./
# Specify method fetch .env!
COPY --from=builder /srv/.env.local ./
CMD ["/srv/app"]