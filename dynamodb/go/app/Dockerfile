FROM golang:1.22

WORKDIR /app

COPY go.mod .
COPY src /app/src

RUN go mod tidy

EXPOSE 8080
