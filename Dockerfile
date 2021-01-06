FROM golang:1.15

# Author info：
# - [ I don't like use MAINTAINER to define the Author]
# - Author : A-Donga
#
# Verwion : build-0.0.1

# Copy file
WORKDIR /go/src/app

COPY . .

ENV GOPROXY https://goproxy.io

RUN go get -d -v ./...
RUN go build -o app cmd/main.go

CMD ["./app"]