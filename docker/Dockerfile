FROM golang:1.23 AS builder

WORKDIR /build

COPY go.mod go.sum ./
COPY impl/ ./impl/

RUN go mod download
RUN go build -o astroapi ./impl/main.go



FROM centos:7

WORKDIR /app

RUN mkdir -p /app/bin /app/conf /app/logs

COPY --from=builder /build/astroapi /app/bin/
COPY ./docker/conf/config.yaml /app/conf/

RUN chmod +x /app/bin/astroapi

EXPOSE 8080
CMD ["/app/bin/astroapi", "-c", "/app/conf/config.yaml"]

