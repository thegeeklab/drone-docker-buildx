FROM docker:20.10-dind@sha256:d20f0866ee1aca6b0eb7bc8bdb0d0c938721c7f79219b1d3bc1aa58666928586

LABEL maintainer="Robert Kaussow <mail@thegeeklab.de>"
LABEL org.opencontainers.image.authors="Robert Kaussow <mail@thegeeklab.de>"
LABEL org.opencontainers.image.title="drone-docker-buildx"
LABEL org.opencontainers.image.url="https://github.com/thegeeklab/drone-docker-buildx"
LABEL org.opencontainers.image.source="https://github.com/thegeeklab/drone-docker-buildx"
LABEL org.opencontainers.image.documentation="https://github.com/thegeeklab/drone-docker-buildx"

ARG BUILDX_VERSION

# renovate: datasource=github-releases depName=docker/buildx
ENV BUILDX_VERSION="${BUILDX_VERSION:-v0.6.3}"

ENV DOCKER_HOST=unix:///var/run/docker.sock

RUN apk --update add --virtual .build-deps curl && \
    mkdir -p /usr/lib/docker/cli-plugins/ && \
    curl -SsL -o /usr/lib/docker/cli-plugins/docker-buildx "https://github.com/docker/buildx/releases/download/v${BUILDX_VERSION##v}/buildx-v${BUILDX_VERSION##v}.linux-amd64" && \
    chmod 755 /usr/lib/docker/cli-plugins/docker-buildx && \
    apk del .build-deps && \
    rm -rf /var/cache/apk/* && \
    rm -rf /tmp/*

ADD release/amd64/drone-docker-buildx /bin/

ENTRYPOINT ["/usr/local/bin/dockerd-entrypoint.sh", "drone-docker-buildx"]
