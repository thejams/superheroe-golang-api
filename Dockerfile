# Start from golang base image
FROM golang:alpine as build

# Install git and make.
# Git is required for fetching the dependencies.
# Make is required for building the project
RUN apk update \
    && apk add --no-cache git \
    && apk add --no-cache make

# Set the current working directory inside the container 
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Install all dependencies. 
RUN go mod vendor 

# Build proyect
RUN make mod && make build

FROM alpine

WORKDIR /app

COPY --from=build /app/build/bin/superheroe-golang-api /app

# Start Mongo DB
RUN make mongo_start && RUN mongo_prepare

# Expose port 5000 in the container
EXPOSE 5000

# Run project
CMD ["./build/bin/superheroe-golang-api"]