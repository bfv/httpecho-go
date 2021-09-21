FROM alpine:latest

RUN addgroup -S httpecho && adduser -S httpecho -G httpecho

RUN mkdir /app
COPY httpecho /app

RUN chown -r httpecho:httpecho /app && \
    chmod +x /app/httpecho

USER httpecho

CMD [ "/app/httpecho" ]
