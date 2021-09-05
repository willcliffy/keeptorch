FROM golang:1.13

RUN mkdir -p /app
WORKDIR /app

ADD . /app

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]
RUN go build  -o /dist/api /app/main.go

EXPOSE 8080

CMD ["/dist/api"]
