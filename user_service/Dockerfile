FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o user-service ./cmd/user-service

EXPOSE 8080 50051

CMD ["./user-service"]