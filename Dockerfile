FROM golang:latest
RUN mkdir app
COPY . /app
WORKDIR /app
RUN go build -o build/main cmd/app/main.go
CMD ["/app/build/main"]
