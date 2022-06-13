FROM golang:alpine as builder

WORKDIR /go/src/go-admin/internal/app/
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="SliverHorn@sliver_horn@qq.com"

WORKDIR /go/src/go-admin/internal/app/

COPY --from=0 /go/src/go-admin/internal/app/server ./
COPY --from=0 /go/src/go-admin/internal/app/resource ./resource/
COPY --from=0 /go/src/go-admin/internal/app/config.docker.yaml ./

EXPOSE 8888
ENTRYPOINT ./server -c config.docker.yaml
