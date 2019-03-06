FROM golang:1.12 as build
WORKDIR /opt/genesis
COPY . .
RUN CGO_ENABLED=0 go build -o genesis cmd/web/*

FROM alpine:latest
WORKDIR /opt/genesis
COPY --from=build /opt/genesis/genesis /opt/genesis/bin/genesis
COPY --from=build /opt/genesis/static /opt/genesis/static
EXPOSE 8080
ENTRYPOINT ["/opt/genesis/bin/genesis"]
