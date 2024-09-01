FROM golang:1.23

RUN apt-get update && apt-get upgrade -y \
    && curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.0/install.sh | bash \
    && /bin/bash -c "source ~/.nvm/nvm.sh && nvm install 20" \
    && /bin/bash -c "source ~/.nvm/nvm.sh && npm install -g nodemon"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o user-service ./cmd/user_service

EXPOSE 8080 50051

CMD ["./user-service"]