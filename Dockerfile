FROM golang:1.17 as builder

ARG SHA

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-X 'main.version="$SHA"'"  -o bin/app

FROM alpine:latest as release
RUN apk update && apk add ca-certificates sqlite && rm -rf /var/cache/apk/*
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY --from=builder /app/bin/app .
CMD ["./app"]