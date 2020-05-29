FROM golang:1.13

WORKDIR /go/src/
COPY ../ .

RUN go get -d -v ./...
RUN go build ./ -o $GOPATH/src/github.com/achwanyusuf/user-management/grpclient
RUN ./grpcclient

CMD ["app"]