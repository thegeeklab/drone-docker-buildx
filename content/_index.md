---
title: drone-docker-buildx
---

[![Build Status](https://img.shields.io/drone/build/thegeeklab/drone-docker-buildx?logo=drone&server=https%3A%2F%2Fdrone.thegeeklab.de)](https://drone.thegeeklab.de/thegeeklab/drone-docker-buildx)
[![Docker Hub](https://img.shields.io/badge/dockerhub-latest-blue.svg?logo=docker&logoColor=white)](https://hub.docker.com/r/thegeeklab/drone-docker-buildx)
[![Quay.io](https://img.shields.io/badge/quay-latest-blue.svg?logo=docker&logoColor=white)](https://quay.io/repository/thegeeklab/drone-docker-buildx)
[![GitHub contributors](https://img.shields.io/github/contributors/thegeeklab/drone-docker-buildx)](https://github.com/thegeeklab/drone-docker-buildx/graphs/contributors)
[![Source: GitHub](https://img.shields.io/badge/source-github-blue.svg?logo=github&logoColor=white)](https://github.com/thegeeklab/drone-docker-buildx)
[![License: MIT](https://img.shields.io/github/license/thegeeklab/drone-docker-buildx)](https://github.com/thegeeklab/drone-docker-buildx/blob/main/LICENSE)

Drone plugin to build and publish multiarch Docker images with buildx.

<!-- prettier-ignore-start -->
<!-- spellchecker-disable -->
{{< toc >}}
<!-- spellchecker-enable -->
<!-- prettier-ignore-end -->

## Versioning

The tags follow the major version of Docker, e.g. `20`, and the minor and patch parts reflect the `version` of the plugin. A full example would be `20.12.5`. Minor versions can introduce breaking changes, while patch versions can be considered non-breaking.

## Usage

{{< hint type=important >}}
Be aware that the this plugin requires [privileged](https://docs.drone.io/pipeline/docker/syntax/steps/#privileged-mode) capabilities, otherwise the integrated Docker daemon is not able to start.
{{< /hint >}}

```YAML
kind: pipeline
name: default

steps:
  - name: docker
    image: thegeeklab/drone-docker-buildx
    privileged: true
    settings:
      username: octocat
      password: secure
      repo: octocat/example
      tags: latest
```

### Parameters

<!-- prettier-ignore-start -->
<!-- spellchecker-disable -->
{{< propertylist name=drone-docker-buildx.data >}}
<!-- spellchecker-enable -->
<!-- prettier-ignore-end -->

## Build

Build the binary with the following command:

```Shell
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GO111MODULE=on

make build
```

Build the Docker image with the following command:

```Shell
docker build --file docker/Dockerfile.amd64 --tag thegeeklab/drone-docker-buildx .
```

## Test

```Shell
docker run --rm \
  -e PLUGIN_TAG=latest \
  -e PLUGIN_REPO=octocat/hello-world \
  -e DRONE_COMMIT_SHA=00000000 \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  --privileged \
  thegeeklab/drone-docker-buildx --dry-run
```
