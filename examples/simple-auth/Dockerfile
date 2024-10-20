# Stage 1: Build the Go binary
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download and cache the Go modules
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go binary for Linux
RUN CGO_ENABLED=0 GOPRIVATE=github.com/maddalax LOG_LEVEL=debug go run github.com/maddalax/htmgo/cli/htmgo@latest build

RUN CGO_ENABLED=1 GOOS=linux go build -tags prod -o ./dist -a -ldflags '-linkmode external -extldflags "-static"' .


# Stage 2: Create the smallest possible image
FROM gcr.io/distroless/base-debian11

# Set the working directory inside the container
WORKDIR /app

# Copy the Go binary from the    builder stage
COPY --from=builder /app/dist .

# Expose the necessary port (replace with your server port)
EXPOSE 3000


# Command to run the binary
CMD ["./simpleauth"]
