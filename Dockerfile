## BUILDER #####################################################################

FROM golang:alpine as builder

WORKDIR /go/src/github.com/essentialkaos/yo

COPY . .

RUN apk add --no-cache git=~2.26 make=4.3-r0 && \
    make deps && \
    make all

## FINAL IMAGE #################################################################

FROM alpine:3.10

LABEL name="Yo Image" \
      vendor="ESSENTIAL KAOS" \
      maintainer="Anton Novojilov" \
      license="Apache-2.0" \
      version="2020.06.13"

COPY --from=builder /go/src/github.com/essentialkaos/yo/yo /usr/bin/

# hadolint ignore=DL3018
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["yo"]

################################################################################
