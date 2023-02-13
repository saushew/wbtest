FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go build -o app -v ./cmd

CMD ["./app"]

