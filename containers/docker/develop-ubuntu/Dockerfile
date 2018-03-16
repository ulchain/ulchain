FROM ubuntu:xenial

ENV PATH=/usr/lib/go-1.9/bin:$PATH

RUN \
  apt-get update && apt-get upgrade -q -y && \
  apt-get install -y --no-install-recommends golang-1.9 git make gcc libc-dev ca-certificates && \
  git clone --depth 1 https://github.com/epvchain/go-epvchain && \
  (cd go-epvchain && make gepv) && \
  cp go-epvchain/build/bin/gepv /gepv && \
  apt-get remove -y golang-1.9 git make gcc libc-dev && apt autoremove -y && apt-get clean && \
  rm -rf /go-epvchain

EXPOSE 7545
EXPOSE 50303

ENTRYPOINT ["/gepv"]
