FROM golang:1.17.8-alpine
WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /docker-gs-ping cmd/gobank/main.go

EXPOSE 8080

CMD [ "/docker-gs-ping" ]