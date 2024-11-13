FROM golang:latest

WORKDIR /saas-microservice

COPY . .

RUN go get -d -v ./...

RUN GOOS=linux GOARCH=amd64 go build -o app ./cmd/app

EXPOSE 8001
CMD ["./app"]