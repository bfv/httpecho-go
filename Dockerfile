FROM alpine:latest

# ubuntu
# RUN useradd httpecho

# alpine
RUN adduser -S httpecho

RUN mkdir /app
COPY build/httpecho /app

RUN chmod +x /app/httpecho

USER httpecho

EXPOSE 1323

CMD [ "/app/httpecho" ]
