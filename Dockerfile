FROM golang:alpine

WORKDIR /app
COPY . /app
ENV GO111MODULE=on
RUN go build -o ./bin/opensocks-cloud ./main.go

ENTRYPOINT ["./bin/opensocks-cloud"]

