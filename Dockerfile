FROM golang:1.20-alpine

LABEL authors="Hamza Hasan"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o airfreight

EXPOSE 9010

CMD ["./airfreight"]