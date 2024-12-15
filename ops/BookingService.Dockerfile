# Booking Service Dockerfile
FROM golang:1.23.4 as builder

WORKDIR /app

# Copy files
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the booking service binary as a statically linked executable
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o booking-service ./cmd/booking_service.go

# Final stage: Minimal image
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/booking-service .

# Ensure execution permissions
RUN chmod +x booking-service

EXPOSE 50052

CMD ["./booking-service"]
