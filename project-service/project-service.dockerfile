FROM alpine:latest

RUN mkdir /app

COPY projectApp /app

CMD ["/app/projectApp"]