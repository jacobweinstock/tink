FROM alpine:3.20.3

ARG TARGETOS
ARG TARGETARCH

RUN apk add --no-cache --update --upgrade ca-certificates=20240705-r0

COPY bin/tink-controller-v1alpha2-${TARGETOS}-${TARGETARCH} /usr/bin/tink-controller

ENTRYPOINT ["/usr/bin/tink-controller"]
