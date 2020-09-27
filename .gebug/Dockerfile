FROM golang:1.14
RUN go get github.com/githubnemo/CompileDaemon
RUN go get github.com/go-delve/delve/cmd/dlv

WORKDIR /src
COPY . .

RUN go build -gcflags="all=-N -l" -o /build/package
ENTRYPOINT dlv --listen=:8888 --headless=true --api-version=2 --accept-multiclient exec /build/package