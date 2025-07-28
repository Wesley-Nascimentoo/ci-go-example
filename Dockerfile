FROM golang:1.20

WORKDIR /app

COPY . .

RUN go mod init example.com/m

RUN go build -o math

CMD ["./math"]