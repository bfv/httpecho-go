FROM ubuntu:latest

# ubuntu
RUN useradd httpecho

RUN mkdir /app
COPY build/httpecho-linux /app/

RUN mv httpecho-linux httpecho && \
    chmod +x /app/httpecho

USER httpecho

EXPOSE 1323

CMD [ "/app/httpecho" ]
