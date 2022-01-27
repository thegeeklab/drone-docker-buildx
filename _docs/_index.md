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

The tags follow the major version of Docker, e.g. `20`, the minor and patch part reflects the "version" of the plugin. A full example would be `20.12.5`. Minor versions may introduce breaking changes, while patch versions may be considered non-breaking.

## Build

Build the binary with the following command:

```Shell
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GO111MODULE=on

go build -v -a -tags netgo -o release/amd64/drone-docker-buildx ./cmd/drone-docker-buildx/
```

Build the Docker image with the following command:

```Shell
docker build --file docker/Dockerfile.amd64 --tag thegeeklab/drone-docker-buildx .
```

## Usage

{{< hint warning >}}
**Note**\
Be aware that the this plugin requires privileged capabilities, otherwise the
integrated Docker daemon is not able to start.
{{< /hint >}}

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

### Parameters

dry_run
: disables docker push

drone_remote_url
: sets the git remote url

mirror
: sets a registry mirror to pull images

storage_driver
: sets the docker daemon storage driver

storage_path
: sets the docker daemon storage path (default `/var/lib/docker`)

bip
: allows the docker daemon to bride ip address

mtu
: sets docker daemon custom mtu setting

custom_dns
: sets custom docker daemon dns server

custom_dns_search
: sets custom docker daemon dns search domain

insecure
: allows the docker daemon to use insecure registries

ipv6
: enables docker daemon ipv6 support

experimental
: enables docker daemon experimental mode

debug
: enables verbose debug mode for the docker daemon

daemon_off
: disables the startup of the docker daemon

buildkit_config
: sets content of the docker buildkit json config

dockerfile
: sets dockerfile to use for the image build (default `./Dockerfile`)

context
: sets the path of the build context to use (default `./`)

tags
: sets repository tags to use for the image; tags can also be loaded from a `.tags` file (default `latest`)

auto_tag
: generates tag names automatically based on git branch and git tag

auto_tag_suffix
: generates tag names with the given suffix

build_args
: sets custom build arguments for the build

build_args_from_env
: forwards environment variables as custom arguments to the build

quiet
: enables suppression of the build output

target
: sets the build target to use

cache_from
: sets images to consider as cache sources

pull_image
: enforces to pull base image at build time (default `true`)

compress
: enables compression of the build context using gzip

output
: sets the [export action](https://docs.docker.com/engine/reference/commandline/buildx_build/#output) for the build result (format: `path` or `type=TYPE[,KEY=VALUE]`)

repo
: sets repository name for the image

registry
: sets docker registry to authenticate with (default `https://index.docker.io/v1/`)

username
: sets username to authenticates with

password
: sets password to authenticates with

email
: sets email address to authenticates with

config
: sets content of the docker daemon json config

purge
: enables cleanup of the docker environment at the end of a build (default `true`)

no_cache
: disables the usage of cached intermediate containers

add_host
: sets additional host:ip mapping

platforms
: sets target platform for build
