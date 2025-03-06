FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY impl/ ./
RUN go build -o astroapi ./main.go

EXPOSE 8080
CMD ["/app/astroapi"]
