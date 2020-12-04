FROM golang:1.15-buster as builder
RUN go get -u github.com/gobuffalo/packr/packr && \
    packr --input pkg

COPY . /build/
WORKDIR /build
RUN go build -o webp-utils .


FROM busybox as webp_downloader
ARG version="0.5.1"
WORKDIR /download
RUN wget  -qO- https://storage.googleapis.com/downloads.webmproject.org/releases/webp/libwebp-${version}-linux-x86-64.tar.gz | tar xz -C /download && \
    mv /download/libwebp-${version}-linux-x86-64/* . && \
    ls -la


FROM ubuntu:20.10
WORKDIR  /etc/webp-utils
COPY .docker/default_configuration.json default.json

COPY --from=builder /build/webp-utils /usr/local/bin/webp-utils
COPY --from=webp_downloader /download/bin/* /usr/local/bin/

WORKDIR /workspace
ENTRYPOINT ["/usr/local/bin/webp-utils"]
