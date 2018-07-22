FROM golang:1.9.4

RUN go get github.com/golang/dep/cmd/dep
RUN go get github.com/tomoyane/gin-performance

WORKDIR /go/src/github.com/tomoyane/gin-performance

ENV GOPATH $GOPATH:/go/src
ENV DB_SOURCE="root:root@tcp(docker.for.mac.localhost:3306)/performance_test?charset=utf8&parseTime=True"

RUN dep ensure
RUN go build

CMD ["gin-performance"]