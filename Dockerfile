# Start from golang base image
FROM golang:alpine

# ENV GO111MODULE=on

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Install all dependencies. 
RUN go mod vendor 

# Expose port 5000 in the container
EXPOSE 5000

# Build the Go app
CMD ["go","run","main.go"]