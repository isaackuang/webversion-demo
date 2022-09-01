FROM golang:1.18.2-alpine3.16 AS builder

ADD src /go/src
WORKDIR /go/src
RUN go get && \
    go install

FROM isaackuang/alpine-base:3.13

COPY config /

COPY --from=builder /go/bin/web /go/bin/
