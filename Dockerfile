FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app

VOLUME /data

RUN go build -o main .
CMD ["/app/main"]
