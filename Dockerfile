## BUILDER #####################################################################

FROM golang:alpine as builder

WORKDIR /go/src/github.com/essentialkaos/yo

COPY . .

ENV GO111MODULE=auto

# hadolint ignore=DL3018
RUN apk add --no-cache git make upx && \
    make deps && \
    make all && \
    upx yo

## FINAL IMAGE #################################################################

FROM essentialkaos/alpine:3.15

LABEL org.opencontainers.image.title="yo" \
      org.opencontainers.image.description="Command-line YAML processor" \
      org.opencontainers.image.vendor="ESSENTIAL KAOS" \
      org.opencontainers.image.authors="Anton Novojilov" \
      org.opencontainers.image.licenses="Apache-2.0" \
      org.opencontainers.image.url="https://kaos.sh/yo" \
      org.opencontainers.image.source="https://github.com/essentialkaos/yo"

COPY --from=builder /go/src/github.com/essentialkaos/yo/yo /usr/bin/

# hadolint ignore=DL3018
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["yo"]

################################################################################
