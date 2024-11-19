FROM golang:1.22.3-alpine

COPY . /app

WORKDIR /app

RUN go build -o main main.go

EXPOSE 8050

CMD ["./main"]