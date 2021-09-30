FROM golang:alpine

WORKDIR /go/src/app
COPY . .
RUN go get .
RUN go install .

ENV BCHTTPD_PORT=8080
ENV BCHTTPD_ROOT=/var/www

EXPOSE 8080/tcp

VOLUME /var/www

CMD ["bchttpd"]