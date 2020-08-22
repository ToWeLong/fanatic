FROM golang:latest AS builder
WORKDIR /go/src
COPY . .
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.io,direct
RUN go build -o fanatic

FROM alpine:latest AS final
MAINTAINER towelong <towelong@qq.com>
WORKDIR /app
COPY --from=builder /go/src/fanatic /app/fanatic
COPY --from=builder /go/src/conf /app/conf
EXPOSE 8081
ENTRYPOINT ["/app/fanatic"]