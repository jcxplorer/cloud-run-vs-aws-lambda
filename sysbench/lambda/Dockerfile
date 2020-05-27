FROM amazonlinux:2.0.20200406.0 AS sysbench-builder

ARG SYSBENCH_VERSION=1.0.20

WORKDIR /root

RUN yum install -y tar gzip make automake libtool pkgconfig libaio-devel which && \
    curl -sL https://github.com/akopytov/sysbench/archive/${SYSBENCH_VERSION}.tar.gz | tar xz && \
    cd sysbench-${SYSBENCH_VERSION} && \
    ./autogen.sh && \
    ./configure --without-mysql && \
    make && \
    cp src/sysbench /usr/local/bin/sysbench

FROM golang:1.14.2

WORKDIR /app
COPY . .
COPY --from=sysbench-builder /usr/local/bin/sysbench /package/sysbench
COPY --from=sysbench-builder /lib64/libaio.so.1.0.1 /package/lib/libaio.so.1

RUN go build -o /package/sysbench-lambda ./lambda
