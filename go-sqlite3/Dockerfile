FROM golang:1.20

WORKDIR /app

# COPY go.mod, go.sum and download the dependencies
COPY ./api-sqlite3/go.* ./

# Run and verify
RUN go mod download && go mod verify

# Copy local code to the container image.
COPY ./api-sqlite3/. ./

# Mod tidy
RUN go mod tidy


# tells Docker that the container listens on specified network ports at runtime
EXPOSE 8080

# RUN go build -v -o /usr/local/bin/app ./...

# CMD ["app"]