FROM golang:1.8-alpine as builder
COPY ./ /go/src/
WORKDIR /go/src/
RUN apk add --no-cache make git \
    && make golang

FROM alpine:latest
MAINTAINER Martin Buchleitner "martin@nitram.at"
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/loxonegoweather /opt/loxonegoweather
EXPOSE 8080
ENTRYPOINT ["/loxonegoweather"]


