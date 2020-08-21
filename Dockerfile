FROM golang:alpine
WORKDIR $GOPATH/src
ADD . $GOPATH/src/godocker
RUN go build -o fanatic
EXPOSE 8081
ENTRYPOINT ["./fanatic"]
