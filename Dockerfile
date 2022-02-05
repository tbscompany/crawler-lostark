FROM golang:1.17 as builder

ARG SHA

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-X 'main.version="$SHA"'"  -o bin/app

FROM alpine:latest as release
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /app/bin/app .
CMD ["./app"]