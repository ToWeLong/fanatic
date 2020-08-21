FROM golang:alpine
WORKDIR $GOPATH/src
ADD . $GOPATH/src/godocker
RUN go build fanatic.go
EXPOSE 8081
ENTRYPOINT ["./fanatic"]
