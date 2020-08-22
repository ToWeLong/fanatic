FROM golang:latest AS builder
WORKDIR /go/src
COPY . .
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.io,direct
RUN CGO_ENABLED=0 go build -o App

FROM alpine AS final
MAINTAINER towelong <towelong@qq.com>
WORKDIR /app
COPY --from=builder /go/src/App /app/App
COPY --from=builder /go/src/conf /app/conf
EXPOSE 8081
ENTRYPOINT ["/app/App"]