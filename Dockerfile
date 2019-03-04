FROM golang:1.12 as build
WORKDIR /opt/genesis
COPY . .
RUN go build

FROM alpine:latest
WORKDIR /
COPY --from=build /opt/genesis/genesis /usr/bin/genesis
EXPOSE 80
# ENTRYPOINT ["/usr/bin/genesis"]
