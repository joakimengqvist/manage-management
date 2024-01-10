FROM alpine:latest

RUN mkdir /app

COPY brokerApp /app

EXPOSE 8080

CMD ["/app/brokerApp"]