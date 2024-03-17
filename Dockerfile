# syntax=docker/dockerfile:1
# A url-shortner microservice in Go packaged into a container image.

FROM golang:1.20


# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /url-shortener

EXPOSE 8080


# Run
CMD ["/url-shortener"]