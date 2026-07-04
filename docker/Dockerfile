ARG IMAGE_GOLANG=1.25.5
ARG IMAGE_ALPINE=3.22

FROM golang:${IMAGE_GOLANG}-alpine${IMAGE_ALPINE} AS build

WORKDIR /app

ARG GIT_COMMIT
ARG VERSION
ARG XDATA
ARG ZDATA

ENV GO111MODULE=on
ENV GONOSUMCHECK=github.com/ssl-pinning/*
ENV GONOSUMDB=github.com/ssl-pinning/*
ENV GOPRIVATE=github.com/ssl-pinning/*
ENV GOFLAGS=-mod=mod

ADD . .

RUN apk add --no-cache git

RUN --mount=type=secret,id=GIT_AUTH_TOKEN \
    TOKEN=$(cat /run/secrets/GIT_AUTH_TOKEN) && \
    git config --global url."https://x-access-token:${TOKEN}@github.com/".insteadOf "https://github.com/" && \
    go build \
        -o build/package/usr/local/sbin/ssl-pinning \
        -ldflags "-w -s -X github.com/ssl-pinning/ssl-pinning/pkg/version.gitCommit=${GIT_COMMIT} -X github.com/ssl-pinning/ssl-pinning/pkg/version.version=${VERSION} -X github.com/ssl-pinning/core.x=${XDATA} -X github.com/ssl-pinning/core.z=${ZDATA}" \
        github.com/ssl-pinning/ssl-pinning

FROM alpine:${IMAGE_ALPINE}

COPY --from=build /app/build/package /

ENTRYPOINT ["/usr/local/sbin/ssl-pinning"]
