FROM alpine:3.7

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 https://github.com/epvchain/go-epvchain && \
  (cd go-epvchain && make gepv) && \
  cp go-epvchain/build/bin/gepv /gepv && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /go-epvchain && rm -rf /var/cache/apk/*

EXPOSE 7545
EXPOSE 50303

ENTRYPOINT ["/gepv"]
