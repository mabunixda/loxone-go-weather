FROM debian:jessie
MAINTAINER Martin Buchleitner "martin@nitram.at"
RUN apt-get update
RUN apt-get install -y ca-certificates

COPY loxonegoweather /opt/loxonegoweather
RUN chmod 755 /opt/loxonegoweather

ENTRYPOINT ["/opt/loxonegoweather"]
EXPOSE 8080
CMD [""]
