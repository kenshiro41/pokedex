FROM golang:latest

RUN mkdir /go/src/go-api

WORKDIR /go/src/go-api

ADD . /go/src/go-api

RUN go get -u github.com/gin-gonic/gin && \
  go get github.com/jinzhu/gorm && \
  go get github.com/jinzhu/gorm/dialects/postgres && \
  go get github.com/pilu/fresh && \
  go get github.com/gin-contrib/cors

RUN apt-get clean all