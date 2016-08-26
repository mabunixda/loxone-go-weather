FROM debian:jessie
MAINTAINER Martin Buchleitner "martin@nitram.at"

COPY loxonegoweather /opt/loxonegoweather
RUN chmod 644 /opt/loxonegoweather

ENTRYPOINT ["/opt/loxonegoweather"]
EXPOSE 8080
CMD [""]