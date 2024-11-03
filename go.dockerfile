FROM golang:1.23.1-alpine3.20

COPY ./app /go/app/

WORKDIR /go/app/

RUN go mod download && go mod tidy
RUN go install github.com/air-verse/air@latest

CMD ["air", "-c", ".air.toml"]

EXPOSE 8080
