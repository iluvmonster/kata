FROM golang

RUN apt-get update -y && apt-get install -y vim git
COPY src /go/src/github.com/iluvmonster
