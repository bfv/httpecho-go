FROM ubuntu:latest

# ubuntu
RUN useradd httpecho

RUN mkdir /app
COPY build/httpecho-linux /app/

RUN mv /app/httpecho-linux /app/httpecho && \
    chmod +x /app/httpecho

USER httpecho

EXPOSE 1323

CMD [ "/app/httpecho" ]
