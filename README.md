# drone-docker-buildx

Drone plugin to build multiarch Docker images with buildx

[![Build Status](https://img.shields.io/drone/build/thegeeklab/drone-docker-buildx?logo=drone&server=https%3A%2F%2Fdrone.thegeeklab.de)](https://drone.thegeeklab.de/thegeeklab/drone-docker-buildx)
[![Docker Hub](https://img.shields.io/badge/dockerhub-latest-blue.svg?logo=docker&logoColor=white)](https://hub.docker.com/r/thegeeklab/drone-docker-buildx)
[![Quay.io](https://img.shields.io/badge/quay-latest-blue.svg?logo=docker&logoColor=white)](https://quay.io/repository/thegeeklab/drone-docker-buildx)
[![Go Report Card](https://goreportcard.com/badge/github.com/thegeeklab/drone-docker-buildx)](https://goreportcard.com/report/github.com/thegeeklab/drone-docker-buildx)
[![GitHub contributors](https://img.shields.io/github/contributors/thegeeklab/drone-docker-buildx)](https://github.com/thegeeklab/drone-docker-buildx/graphs/contributors)
[![Source: GitHub](https://img.shields.io/badge/source-github-blue.svg?logo=github&logoColor=white)](https://github.com/thegeeklab/drone-docker-buildx)
[![License: MIT](https://img.shields.io/github/license/thegeeklab/drone-docker-buildx)](https://github.com/thegeeklab/drone-docker-buildx/blob/main/LICENSE)

Drone plugin to build multiarch Docker images with buildx. This plugin is a fork of [drone-plugins/drone-docker](https://github.com/drone-plugins/drone-docker).

## Docker Tags

Tags are following the main Docker version e.g. `20.10`, the second part is reflecting the plugin "version". A full example would be `20.10.5`.

## Build

Build the binary with the following command:

```Shell
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GO111MODULE=on

go build -v -a -tags netgo -o release/drone-docker-buildx
```

Build the Docker image with the following command:

```Shell
docker build --file docker/Dockerfile.amd64 --tag thegeeklab/drone-docker-buildx .
```

## Usage

> Notice: Be aware that the tis plugin requires privileged capabilities, otherwise the integrated Docker daemon is not able to start.

```console
docker run --rm \
  -e PLUGIN_TAG=latest \
  -e PLUGIN_REPO=octocat/hello-world \
  -e DRONE_COMMIT_SHA=00000000 \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  --privileged \
  thegeeklab/drone-docker-buildx --dry-run
```

## Contributors

Special thanks goes to all [contributors](https://github.com/thegeeklab/drone-docker-buildx/graphs/contributors). If you would like to contribute,
please see the [instructions](https://github.com/thegeeklab/drone-docker-buildx/blob/main/CONTRIBUTING.md).

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/thegeeklab/drone-docker-buildx/blob/main/LICENSE) file for details.
