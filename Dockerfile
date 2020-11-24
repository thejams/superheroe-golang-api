# Start from golang base image
FROM golang:alpine

# ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Steven Victor <chikodi543@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 
WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Expose port 5000 in the container
EXPOSE 5000

# Build the Go app
CMD ["go","run","main.go"]