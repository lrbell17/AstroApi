    FROM golang:1.23

    WORKDIR /app
    COPY . ./

    RUN go mod download
    RUN go build -o astroapi ./impl/main.go

    EXPOSE 8080
    CMD ["/app/astroapi"]
