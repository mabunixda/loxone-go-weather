FROM alpine:latest
MAINTAINER Martin Buchleitner "martin@nitram.at"

RUN apk --no-cache add ca-certificates

COPY loxonegoweather /opt/loxonegoweather
RUN chmod 755 /opt/loxonegoweather

EXPOSE 8080

CMD ["/opt/loxonegoweather"]
