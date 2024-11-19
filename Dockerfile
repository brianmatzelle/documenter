FROM golang:1.22.4-alpine

# Add necessary build dependencies
RUN apk add --no-cache git

WORKDIR /app

# Copy go mod files first for better layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build with optimizations
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

# Use multi-stage build for smaller final image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=0 /app/main .

EXPOSE ${PORT}

CMD ["./main"]