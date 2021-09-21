FROM ubuntu:latest

RUN useradd httpecho

RUN mkdir /app
COPY build/httpecho /app

RUN chmod +x /app/httpecho

USER httpecho

EXPOSE 1323

CMD [ "/app/httpecho" ]
