# syntax=docker/dockerfile:1.4
FROM golang:1.24 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/bank

FROM alpine:latest

COPY --from=build /bin/bank /bank
EXPOSE 8080
CMD ["/bank"]
