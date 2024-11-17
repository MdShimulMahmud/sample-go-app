FROM golang:1.23.3-alpine3.20 as base

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
COPY *.go ./


# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /sampleapp

EXPOSE 8080

# Run
CMD ["/sampleapp"]

