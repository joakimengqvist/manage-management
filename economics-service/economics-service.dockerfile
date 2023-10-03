FROM alpine:latest

RUN mkdir /app

COPY economicsApp /app

CMD ["/app/economicsApp"]