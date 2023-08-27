FROM golang:1.20-alpine

WORKDIR /app

ENV CGO_ENABLED=1

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
