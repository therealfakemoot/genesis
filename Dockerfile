FROM golang:1.12-alpine as builder
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

WORKDIR /opt/genesis

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o genesis cmd/web/*

FROM alpine:latest
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /opt/genesis/genesis /opt/genesis/genesis
COPY --from=builder /opt/genesis/static /opt/genesis/static

WORKDIR /opt/genesis

EXPOSE 8888
ENTRYPOINT ["/opt/genesis/genesis"]
