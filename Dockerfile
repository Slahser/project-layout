# syntax = docker/dockerfile:experimental
FROM golang:1.16.4-alpine AS build_base

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk add --no-cache ca-certificates upx git tzdata libc6-compat make build-base
RUN ln -snf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" > /etc/timezone

ENV TZ=Asia/Shanghai
ENV GOPROXY https://goproxy.cn,https://mirrors.aliyun.com/goproxy,direct
ENV GO111MODULE on

ARG GITLAB_LOGIN=""
ARG GITLAB_TOKEN=""

RUN echo "machine gitlab.com login ${GITLAB_LOGIN} password ${GITLAB_TOKEN}" > /root/.netrc
RUN chmod 600 /root/.netrc

WORKDIR /src

#ENV CGO_ENABLED=0
COPY go.mod .
COPY go.sum .

RUN go mod download -x
RUN go get github.com/go-delve/delve/cmd/dlv

COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build \
    GOOS=linux GOARCH=amd64 \
#    日常debug 保留此行
#    go build -gcflags "all=-N -l" \
#    线上优化代码,保留此行
    go build -ldflags="-s -w" \
    -o ./out/ttapp ./cmd/apiserver
#RUN upx --best ./out/ttapp -o ./out/ttapp_min

FROM build_base AS gotest
RUN --mount=type=cache,target=/root/.cache/go-build \
    go test -v
#RUN go test -v

#FROM scratch AS bin_base
FROM alpine:20200917 AS bin_base

COPY --from=build_base /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=build_base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build_base /src/out/ttapp /ttapp
COPY --from=build_base /go/bin/dlv /dlv

EXPOSE 8080 8888

#线上优化代码,保留此行
CMD ["/ttapp"]
#日常debug 保留此行
#CMD ["/dlv", "--listen=:8888", "--headless=true", "--api-version=2", "--log","exec", "/ttapp"]