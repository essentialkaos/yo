## BUILDER #####################################################################

FROM golang:alpine as builder

WORKDIR /go/src/github.com/essentialkaos/yo

COPY . .

ENV GO111MODULE=auto

RUN apk add --no-cache git=~2.30 make=4.3-r0 upx=3.96-r1 && \
    make deps && \
    make all && \
    upx yo

## FINAL IMAGE #################################################################

FROM alpine:3.13

LABEL name="Yo Image" \
      vendor="ESSENTIAL KAOS" \
      maintainer="Anton Novojilov" \
      license="Apache-2.0" \
      version="2021.05.01"

COPY --from=builder /go/src/github.com/essentialkaos/yo/yo /usr/bin/

# hadolint ignore=DL3018
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["yo"]

################################################################################
