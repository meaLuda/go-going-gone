FROM golang:1.20

WORKDIR /app

# COPY go.mod, go.sum and download the dependencies
COPY ./dairy-api/go.* ./

# Run and verify
RUN go mod download && go mod verify

# Copy local code to the container image.
COPY ./dairy-api/. ./

# Mod tidy final check incase somethings missing.
RUN go mod tidy


# tells Docker that the container listens on specified network ports at runtime
EXPOSE 8080
