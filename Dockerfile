FROM golang:1.23.1-alpine3.20

WORKDIR /app
COPY . .

RUN go mod download && go mod tidy
RUN go build -o main ./cmd/main.go

EXPOSE $PORT

CMD ["/app/main"]
