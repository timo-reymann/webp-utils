FROM circleci/golang:1.17 as builder
USER root
WORKDIR /build
COPY . .
RUN make build-docker-binary


FROM busybox as webp_downloader
ARG version="0.5.1"
WORKDIR /download
RUN wget  -qO- https://storage.googleapis.com/downloads.webmproject.org/releases/webp/libwebp-${version}-linux-x86-64.tar.gz | tar xz -C /download && \
    mv /download/libwebp-${version}-linux-x86-64/* . && \
    ls -la


FROM ubuntu:22.04
WORKDIR  /etc/webp-utils
COPY .docker/configurations/* .

COPY --from=builder /build/webp-utils /usr/local/bin/webp-utils
COPY --from=webp_downloader /download/bin/* /usr/local/bin/

WORKDIR /workspace
ENTRYPOINT ["/usr/local/bin/webp-utils"]
