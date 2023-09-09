FROM golang:1.21.1-alpine3.18

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o test

CMD [ "go", "run", "./main.go" ]