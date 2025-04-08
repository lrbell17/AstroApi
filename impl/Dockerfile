FROM golang:1.23 AS builder

WORKDIR /build

COPY impl/ .

RUN go mod download
RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o astroapi ./cmd/main.go


FROM centos:7

WORKDIR /app

RUN mkdir -p /app/bin /app/conf /app/logs /app/data

COPY --from=builder /build/astroapi /app/bin/
COPY ./docker/conf/config.yaml /app/conf/
COPY ./docker/conf/private.pem /app/conf/
COPY ./docker/conf/jwk.json /app/conf/
COPY ./docker/data/ /app/data/

RUN chmod +x /app/bin/astroapi

EXPOSE 8080
CMD ["/bin/sh", "-c", "/app/bin/astroapi -c /app/conf/config.yaml & while [ ! -f /app/logs/astroapi.log ]; do sleep 1; done; tail -f /app/logs/astroapi.log"]
