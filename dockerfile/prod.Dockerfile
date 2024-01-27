FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary

EXPOSE 8011

ENTRYPOINT ["/usr/local/bin/uwang-rest-storage"]