# Building server application
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY ["app/go.mod", "app/go.sum", "./"]

RUN go mod download

COPY app ./

RUN go build -o ./build ./cmd/main.go

# Running server
FROM alpine

COPY --from=builder /app/build/main /
# Copy .env to the server dir
# COPY .env .

EXPOSE 8081

CMD ["/main"]