FROM golang:1.19-alpine3.18 as builder

WORKDIR /app
COPY . /app

WORKDIR /app/src

RUN go build \
    -ldflags "-X app/version.GitCommit=`git rev-parse --short=8 HEAD`" \
    -o /build/app

FROM alpine:3.18

RUN apk add --no-cache \
    bash \
    make \
    tzdata

WORKDIR /

COPY --from=builder /build/app /simple-core-bank
COPY makefile .

ENTRYPOINT [ "/simple-core-bank" ]