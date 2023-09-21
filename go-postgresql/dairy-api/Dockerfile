FROM golang:1.20


RUN apt-get update \
 && apt-get install

WORKDIR /app

# # COPY go.mod, go.sum and download the dependencies
COPY ./go.* .
RUN go mod download

# Copy local code to the container image.
COPY . .

# Mod tidy final check incase somethings missing.
RUN go mod tidy


# tells Docker that the container listens on specified network ports at runtime
EXPOSE 8080
