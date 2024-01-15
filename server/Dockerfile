FROM golang:1.21.6-alpine

LABEL maintainer="Marco Braun"

RUN apk update && apk add --no-cache git bash build-base

# Set the working directory to /app
WORKDIR /app

# Copy the contents of the server directory into /app
COPY . .
COPY .env /app/cmd/.env


# Change directory to where main.go is located
WORKDIR /app/cmd

# Fetch dependencies
RUN go get -d -v ./...

# Build the Go app
RUN go build -o /app/cmd/main .

EXPOSE 8081

CMD ["/app/cmd/main"]
