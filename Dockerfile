FROM golang:1.19

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o /app_bin ./cmd/main.go

EXPOSE 8800

CMD ["/app_bin"]