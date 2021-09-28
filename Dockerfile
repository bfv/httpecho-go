FROM ubuntu:latest

# ubuntu
RUN useradd httpecho

RUN mkdir /app
COPY build/httpecho-linux /app/httpecho

RUN chmod +x /app/httpecho

USER httpecho

EXPOSE 1323

CMD [ "/app/httpecho" ]
