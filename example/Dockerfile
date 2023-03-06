FROM golang:1.19-alpine

RUN mkdir /app 
COPY . /app/
WORKDIR /app

RUN go mod download

RUN go build -o main main.go
CMD ["/app/main"]