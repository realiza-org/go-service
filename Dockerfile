# Stage 1: Build the Go application
FROM golang:alpine as builder

WORKDIR /usr/src/app

# Pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the entire source code
COPY . .

# Build the binary, specifying the entry point as cmd/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /usr/local/bin/app cmd/main.go

# Stage 2: Create a minimal Docker image
FROM scratch
COPY --from=builder /usr/local/bin/app /usr/local/bin/app
CMD ["/usr/local/bin/app"]