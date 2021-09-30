FROM golang:alpine

WORKDIR /app
COPY bchttpd .

ENV BCHTTPD_PORT=8080
ENV BCHTTPD_ROOT=/var/www

EXPOSE 8080/tcp

VOLUME /var/www

CMD ["bchttpd"]