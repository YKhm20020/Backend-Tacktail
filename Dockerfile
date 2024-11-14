FROM golang:1.23.1-alpine3.20

COPY ./app /go/app/

WORKDIR /go/app/
COPY . .

RUN go mod download && go mod tidy
RUN go build -o main .

EXPOSE 8080

CMD ["./main.go"]
