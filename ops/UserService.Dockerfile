# User Service Dockerfile
FROM golang:1.23.4 as builder

WORKDIR /app

# Copy files
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the user service binary as a statically linked executable
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user-service ./cmd/user_service.go

# Final stage: Minimal image
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/user-service .

# Ensure execution permissions
RUN chmod +x user-service

EXPOSE 50051

CMD ["./user-service", "50051", "false"]
