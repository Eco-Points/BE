FROM golang:1.22-alpine

RUN apk add --no-cache git

RUN mkdir /app

WORKDIR /app

COPY ./ /app

RUN go mod tidy

RUN go build -o eco_point

CMD ["./eco_point"]