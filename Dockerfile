FROM golang:1.25-alpine
WORKDIR /app
COPY . .
RUN go test ./...
RUN go build -o service .
EXPOSE 8080
CMD ["./service"]
