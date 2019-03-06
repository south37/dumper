FROM golang:1.12.0-stretch

COPY .  /app

WORKDIR /app

RUN make

EXPOSE 8080

CMD ["/app/dumper"]
