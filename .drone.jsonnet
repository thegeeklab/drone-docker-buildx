local PipelineTest = {
  kind: 'pipeline',
  image_pull_secrets: ['docker_config'],
  name: 'test',
  platform: {
    os: 'linux',
    arch: 'amd64',
  },
  steps: [
    {
      name: 'staticcheck',
      image: 'golang:1.15',
      commands: [
        'go run honnef.co/go/tools/cmd/staticcheck ./...',
      ],
      volumes: [
        {
          name: 'gopath',
          path: '/go',
        },
      ],
    },
    {
      name: 'lint',
      image: 'golang:1.15',
      commands: [
        'go run golang.org/x/lint/golint -set_exit_status ./...',
      ],
      volumes: [
        {
          name: 'gopath',
          path: '/go',
        },
      ],
    },
    {
      name: 'vet',
      image: 'golang:1.15',
      commands: [
        'go vet ./...',
      ],
      volumes: [
        {
          name: 'gopath',
          path: '/go',
        },
      ],
    },
    {
      name: 'test',
      image: 'golang:1.15',
      commands: [
        'go test -cover ./...',
      ],
      volumes: [
        {
          name: 'gopath',
          path: '/go',
        },
      ],
    },
  ],
  volumes: [
    {
      name: 'gopath',
      temp: {},
    },
  ],
  trigger: {
    ref: ['refs/heads/main', 'refs/tags/**', 'refs/pull/**'],
  },
};


local PipelineBuildBinaries = {
  kind: 'pipeline',
  image_pull_secrets: ['docker_config'],
  name: 'build-binaries',
  platform: {
    os: 'linux',
    arch: 'amd64',
  },
  steps: [
    {
      name: 'build',
      image: 'techknowlogick/xgo:go-1.15.x',
      commands: [
        '[ -z "${DRONE_TAG}" ] && BUILD_VERSION=${DRONE_COMMIT_SHA:0:8} || BUILD_VERSION=${DRONE_TAG##v}',
        'mkdir -p release/',
        "cd cmd/drone-docker-buildx && xgo -ldflags \"-s -w -X main.version=$BUILD_VERSION\" -tags netgo -targets 'linux/amd64,linux/arm-6,linux/arm64' -out drone-docker-buildx .",
        'mv /build/* /drone/src/release/',
      ],
    },
    {
      name: 'executable',
      image: 'alpine',
      commands: [
        '$(find release/ -executable -type f | grep drone-docker-buildx-linux-amd64) --help',
      ],
    },
    {
      name: 'compress',
      image: 'alpine',
      commands: [
        'apk add upx',
        'find release/ -maxdepth 1 -executable -type f -exec upx {} \\;',
        'ls -lh release/',
      ],
    },
    {
      name: 'checksum',
      image: 'alpine',
      commands: [
        'cd release/ && sha256sum * > sha256sum.txt',
      ],
    },
    {
      name: 'publish',
      image: 'plugins/github-release',
      settings: {
        overwrite: true,
        api_key: {
          from_secret: 'github_token',
        },
        files: ['release/*'],
        title: '${DRONE_TAG}',
        note: 'CHANGELOG.md',
      },
      when: {
        ref: [
          'refs/tags/**',
        ],
      },
    },
  ],
  depends_on: [
    'test',
  ],
  trigger: {
    ref: ['refs/heads/main', 'refs/tags/**', 'refs/pull/**'],
  },
};

local PipelineBuildContainer(arch='amd64') = {
  kind: 'pipeline',
  image_pull_secrets: ['docker_config'],
  name: 'build-container-' + arch,
  platform: {
    os: 'linux',
    arch: arch,
  },
  steps: [
    {
      name: 'build',
      image: 'golang:1.15',
      commands: [
        '[ -z "${DRONE_TAG}" ] && BUILD_VERSION=${DRONE_COMMIT_SHA:0:8} || BUILD_VERSION=${DRONE_TAG##v}',
        'go build -v -ldflags "-X main.version=$BUILD_VERSION" -a -tags netgo -o release/' + arch + '/drone-docker-buildx ./cmd/drone-docker-buildx',
      ],
    },
    {
      name: 'dryrun',
      image: 'plugins/docker:19',
      settings: {
        config: { from_secret: 'docker_config' },
        dry_run: true,
        dockerfile: 'docker/Dockerfile.' + arch,
        repo: 'thegeeklab/${DRONE_REPO_NAME}',
        username: { from_secret: 'docker_username' },
        password: { from_secret: 'docker_password' },
      },
      depends_on: ['build'],
      when: {
        ref: ['refs/pull/**'],
      },
    },
    {
      name: 'publish-dockerhub',
      image: 'plugins/docker:19',
      settings: {
        config: { from_secret: 'docker_config' },
        auto_tag: true,
        auto_tag_suffix: arch,
        dockerfile: 'docker/Dockerfile.' + arch,
        repo: 'thegeeklab/${DRONE_REPO_NAME}',
        username: { from_secret: 'docker_username' },
        password: { from_secret: 'docker_password' },
      },
      when: {
        ref: ['refs/heads/main', 'refs/tags/**'],
      },
      depends_on: ['dryrun'],
    },
    {
      name: 'publish-quay',
      image: 'plugins/docker:19',
      settings: {
        config: { from_secret: 'docker_config' },
        auto_tag: true,
        auto_tag_suffix: arch,
        dockerfile: 'docker/Dockerfile.' + arch,
        registry: 'quay.io',
        repo: 'quay.io/thegeeklab/${DRONE_REPO_NAME}',
        username: { from_secret: 'quay_username' },
        password: { from_secret: 'quay_password' },
      },
      when: {
        ref: ['refs/heads/main', 'refs/tags/**'],
      },
      depends_on: ['dryrun'],
    },
  ],
  depends_on: [
    'test',
  ],
  trigger: {
    ref: ['refs/heads/main', 'refs/tags/**', 'refs/pull/**'],
  },
};

local PipelineNotifications = {
  kind: 'pipeline',
  image_pull_secrets: ['docker_config'],
  name: 'notifications',
  platform: {
    os: 'linux',
    arch: 'amd64',
  },
  steps: [
    {
      image: 'plugins/manifest',
      name: 'manifest-dockerhub',
      settings: {
        ignore_missing: true,
        auto_tag: true,
        username: { from_secret: 'docker_username' },
        password: { from_secret: 'docker_password' },
        spec: 'docker/manifest.tmpl',
      },
      when: {
        status: ['success'],
      },
    },
    {
      image: 'plugins/manifest',
      name: 'manifest-quay',
      settings: {
        ignore_missing: true,
        auto_tag: true,
        username: { from_secret: 'quay_username' },
        password: { from_secret: 'quay_password' },
        spec: 'docker/manifest-quay.tmpl',
      },
      when: {
        status: ['success'],
      },
    },
    {
      name: 'pushrm-dockerhub',
      image: 'chko/docker-pushrm:1',
      environment: {
        DOCKER_PASS: {
          from_secret: 'docker_password',
        },
        DOCKER_USER: {
          from_secret: 'docker_username',
        },
        PUSHRM_FILE: 'README.md',
        PUSHRM_SHORT: 'Drone plugin to build multiarch Docker images with buildx',
        PUSHRM_TARGET: 'thegeeklab/${DRONE_REPO_NAME}',
      },
      when: {
        status: ['success'],
      },
    },
    {
      name: 'pushrm-quay',
      image: 'chko/docker-pushrm:1',
      environment: {
        APIKEY__QUAY_IO: {
          from_secret: 'quay_token',
        },
        PUSHRM_FILE: 'README.md',
        PUSHRM_TARGET: 'quay.io/thegeeklab/${DRONE_REPO_NAME}',
      },
      when: {
        status: ['success'],
      },
    },
    {
      name: 'matrix',
      image: 'plugins/matrix',
      settings: {
        homeserver: { from_secret: 'matrix_homeserver' },
        roomid: { from_secret: 'matrix_roomid' },
        template: 'Status: **{{ build.status }}**<br/> Build: [{{ repo.Owner }}/{{ repo.Name }}]({{ build.link }}) ({{ build.branch }}) by {{ build.author }}<br/> Message: {{ build.message }}',
        username: { from_secret: 'matrix_username' },
        password: { from_secret: 'matrix_password' },
      },
      when: {
        status: ['success', 'failure'],
      },
    },
  ],
  depends_on: [
    'build-binaries',
    'build-container-amd64',
    'build-container-arm64',
    'build-container-arm',
  ],
  trigger: {
    ref: ['refs/heads/main', 'refs/tags/**'],
    status: ['success', 'failure'],
  },
};

[
  PipelineTest,
  PipelineBuildBinaries,
  PipelineBuildContainer(arch='amd64'),
  PipelineBuildContainer(arch='arm64'),
  PipelineBuildContainer(arch='arm'),
  PipelineNotifications,
]
