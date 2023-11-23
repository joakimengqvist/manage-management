FROM alpine:latest

RUN mkdir /app

COPY invoiceApp /app

CMD ["/app/invoiceApp"]