FROM alpine:latest

RUN mkdir /app

COPY externalCompanyApp /app

CMD ["/app/externalCompanyApp"]