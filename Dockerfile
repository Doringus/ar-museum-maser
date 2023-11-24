FROM golang:latest

WORKDIR /armuseum

COPY . .
RUN go mod tidy