# syntax = docker/dockerfile:experimental
FROM golang:1.15-alpine AS build_base

WORKDIR /Users/slahser/Desktop/tt

ENV CGO_ENABLED=0
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build \
    GOOS=linux GOARCH=amd64 go build -o ./out/ttapp .

FROM build_base AS gotest
RUN --mount=type=cache,target=/root/.cache/go-build \
    go test -v


FROM scratch AS bin-unix

COPY --from=build_base /Users/slahser/Desktop/tt/out/ttapp /ttapp

EXPOSE 8080

# DOCKER_BUILDKIT=1 docker build -t ttapp:multistage .
# https://github.com/chris-crone/containerized-go-dev/blob/main/Dockerfile
CMD ["/ttapp"]