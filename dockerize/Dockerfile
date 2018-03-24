FROM golang:1.10

ADD . /go/src/hello

RUN go install hello
ENV NAME Sachin

ENTRYPOINT /go/bin/hello

EXPOSE 9000

