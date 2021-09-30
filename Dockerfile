FROM golang:alpine

WORKDIR /go/src/app
COPY . .
RUN go get -v ./...
RUN go install -v ./...

ENV BCHTTPD_PORT=8080
ENV BCHTTPD_ROOT=/var/www

EXPOSE 8080/tcp

VOLUME /var/www

CMD ["bchttpd"]