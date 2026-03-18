FROM golang:1.16.3

WORKDIR /tests
COPY . .
RUN go mod download
RUN go test ./...