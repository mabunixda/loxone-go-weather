FROM alpine:latest
MAINTAINER Martin Buchleitner "martin@nitram.at"

RUN apk --no-cache add ca-certificates
ADD server /opt/loxonegoweather
RUN chmod 755 /opt/*

EXPOSE 8080

WORKDIR /opt
SHELL ["/bin/sh"]
ENTRYPOINT ["/opt/loxonegoweather"]


