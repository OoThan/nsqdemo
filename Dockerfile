FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build ./cmd/consumer
CMD ["/app/consumer"]
