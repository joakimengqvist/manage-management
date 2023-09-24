FROM alpine:latest

RUN mkdir /app

COPY notesApp /app

CMD ["/app/notesApp"]