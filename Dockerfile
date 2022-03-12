FROM golang:1.17.8-alpine
WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

ENTRYPOINT ["go", "run", "cmd/gobank/main.go"]