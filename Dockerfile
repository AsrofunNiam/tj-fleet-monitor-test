FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install git jika diperlukan untuk download dependency
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build binary dengan nama 'main-app'
RUN CGO_ENABLED=0 GOOS=linux go build -o /main-app main.go

# ----------------------------
# Tahap 2: Final Image
FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

# Copy binary dari builder
COPY --from=builder /main-app .

# COPY folder configuration secara utuh (termasuk .env)
COPY --from=builder /app/configuration ./configuration

EXPOSE 8080

# Jalankan aplikasi
CMD ["./main-app"]