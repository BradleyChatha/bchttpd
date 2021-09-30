FROM golang

WORKDIR /go/src/bchttpd
COPY . .
RUN sudo mkdir /sys/fs/cgroup/systemd
RUN sudo mount -t cgroup -o none,name=systemd cgroup /sys/fs/cgroup/systemd
RUN go get .
RUN go install .

ENV BCHTTPD_PORT=8080
ENV BCHTTPD_ROOT=/var/www

EXPOSE 8080/tcp

VOLUME /var/www

CMD ["bchttpd"]