# Gunakan image dasar Go
FROM golang:1.23 AS builder

# Set working directory
WORKDIR /app

# Copy go mod dan sum file
COPY go.mod go.sum ./

# Download dependensi
RUN go mod download

# Copy source code
COPY . .

# Build aplikasi
RUN go build -o app . && chmod +x app

# Gunakan image minimal untuk menjalankan aplikasi
FROM alpine:latest

# Install libraries yang diperlukan (jika ada)
RUN apk add --no-cache libc6-compat

# Set working directory
WORKDIR /root/

# Copy binary dari tahap build
COPY --from=builder /app/app .

# Expose port yang digunakan aplikasi
EXPOSE 8080

# Jalankan aplikasi
CMD ["./app"]
