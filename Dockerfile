FROM golang:1.21 as builder

WORKDIR /app

# Copy the go.mod files first to leverage Docker cache
COPY go.mod ./

# Download the Go modules (including the one hosted on GitHub).
RUN go mod download

# Copy the rest of the application's source code.
COPY . .

# Build the application.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o fme .

# Use a Docker multi-stage build to create a lean production image.
FROM alpine:latest


# Set the working directory to /root/
WORKDIR /root/

# Copy the pre-built binary file from the previous stage.
COPY --from=builder /app/fme .

RUN mkdir data
# Run the application.
CMD ["./fme"]
