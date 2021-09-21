FROM alpine:latest

RUN addgroup -S httpecho && adduser -S httpecho -G httpecho

RUN mkdir /app
COPY build/httpecho /app

RUN chmod +x /app/httpecho

USER httpecho

CMD [ "/app/httpecho" ]
