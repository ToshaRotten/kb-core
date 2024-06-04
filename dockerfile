FROM golang:latest

WORKDIR /app

COPY . .
RUN go mod download

COPY *.go ./

run go build .

EXPOSE 8081

CMD ["./main"]
