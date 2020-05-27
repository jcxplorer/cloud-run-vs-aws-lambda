FROM golang:1.14.2 AS go-builder
WORKDIR /app
COPY . .
RUN go build -o /package/sysbench-cloud-run ./cloud-run

FROM debian:buster-slim
ARG SYSBENCH_VERSION=1.0.20
WORKDIR /root
RUN apt-get update && \
    apt-get install -y curl build-essential automake autoconf pkg-config libaio-dev libtool && \
    curl -sL https://github.com/akopytov/sysbench/archive/${SYSBENCH_VERSION}.tar.gz | tar xz && \
    cd sysbench-${SYSBENCH_VERSION} && \
    ./autogen.sh && \
    ./configure --without-mysql && \
    make && \
    cp src/sysbench /usr/local/bin/sysbench

COPY --from=go-builder /package/sysbench-cloud-run /usr/local/bin/sysbench-cloud-run
CMD ["/usr/local/bin/sysbench-cloud-run"]
EXPOSE 8080
