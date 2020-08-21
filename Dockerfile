
FROM golang:alpine AS development
WORKDIR $GOPATH/src
COPY . .
RUN go build -o fanatic fanatic.go

FROM alpine:latest AS production
WORKDIR /root/
COPY --from=development /go/src/fanatic .
EXPOSE 8081
ENTRYPOINT ["./fanatic"]

