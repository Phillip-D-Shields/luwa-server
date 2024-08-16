FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN go build -o luwaServer .

# minimal image for deployment
FROM alpine:latest

RUN mkdir /app

# Copy the binary and other necessary files from the builder stage
COPY --from=builder /app/luwaServer .

COPY views /views

EXPOSE 8080

CMD ["./luwaServer"]