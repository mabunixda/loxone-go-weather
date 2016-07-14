FROM debian:jessie
MAINTAINER Martin Buchleitner "martin@nitram.at"

COPY LoxoneGoCalendar /opt/LoxoneGoCalendar
RUN chmod 644 /opt/LoxoneGoCalendar

ENTRYPOINT ["/opt/LoxoneGoCalendar"]
EXPOSE 8080
CMD [""]