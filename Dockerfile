FROM golang:1.19.3-alpine AS build

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . ./

RUN go build -o /blog

FROM scratch

WORKDIR /

COPY --from=build /blog /blog

EXPOSE 8080

ENTRYPOINT ["/blog"]
