FROM golang:1.19.3-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN go build -o /blog

EXPOSE 8080

CMD ["/blog"]
