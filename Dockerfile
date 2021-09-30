FROM golang:alpine AS BUILD

WORKDIR /build
COPY . .
RUN go get .
RUN go build . -o bchttpd

FROM alpine:latest AS RUN

COPY --from=BUILD /build /app
WORKDIR /app

ENV BCHTTPD_PORT=8080
ENV BCHTTPD_ROOT=/var/www

EXPOSE 8080/tcp

VOLUME /var/www

CMD ["bchttpd"]