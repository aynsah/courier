# Start from the latest golang base image
FROM golang:latest as builder

# Add Maintainer Info
LABEL maintainer="Nathieq Syah <nathieqs16@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod vendor

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN cd services/user && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


######## Start a new stage from scratch #######
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy configuration File
RUN mkdir config
COPY --from=builder /app/services/user/config/config.json ./config

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/services/user/main .

# Expose port 8801 to the outside world
EXPOSE 8801

# Command to run the executable
CMD ["./main"]