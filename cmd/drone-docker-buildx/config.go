package main

import (
	"github.com/thegeeklab/drone-docker-buildx/plugin"
	"github.com/urfave/cli/v2"
)

// settingsFlags has the cli.Flags for the plugin.Settings.
func settingsFlags(settings *plugin.Settings) []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "dry-run",
			EnvVars:     []string{"PLUGIN_DRY_RUN"},
			Usage:       "dry run disables docker push",
			Destination: &settings.Dryrun,
		},
		&cli.StringFlag{
			Name:        "remote.url",
			EnvVars:     []string{"DRONE_REMOTE_URL"},
			Usage:       "git remote url",
			Destination: &settings.Build.Remote,
		},
		&cli.StringFlag{
			Name:        "daemon.mirror",
			EnvVars:     []string{"PLUGIN_MIRROR", "DOCKER_PLUGIN_MIRROR"},
			Usage:       "docker daemon registry mirror",
			Destination: &settings.Daemon.Mirror,
		},
		&cli.StringFlag{
			Name:        "daemon.storage-driver",
			EnvVars:     []string{"PLUGIN_STORAGE_DRIVER"},
			Usage:       "docker daemon storage driver",
			Destination: &settings.Daemon.StorageDriver,
		},
		&cli.StringFlag{
			Name:        "daemon.storage-path",
			EnvVars:     []string{"PLUGIN_STORAGE_PATH"},
			Usage:       "docker daemon storage path",
			Value:       "/var/lib/docker",
			Destination: &settings.Daemon.StoragePath,
		},
		&cli.StringFlag{
			Name:        "daemon.bip",
			EnvVars:     []string{"PLUGIN_BIP"},
			Usage:       "docker daemon bride ip address",
			Destination: &settings.Daemon.Bip,
		},
		&cli.StringFlag{
			Name:        "daemon.mtu",
			EnvVars:     []string{"PLUGIN_MTU"},
			Usage:       "docker daemon custom mtu setting",
			Destination: &settings.Daemon.MTU,
		},
		&cli.StringSliceFlag{
			Name:        "daemon.dns",
			EnvVars:     []string{"PLUGIN_CUSTOM_DNS"},
			Usage:       "docker daemon dns server",
			Destination: &settings.Daemon.DNS,
		},
		&cli.StringSliceFlag{
			Name:        "daemon.dns-search",
			EnvVars:     []string{"PLUGIN_CUSTOM_DNS_SEARCH"},
			Usage:       "docker daemon dns search domains",
			Destination: &settings.Daemon.DNSSearch,
		},
		&cli.BoolFlag{
			Name:        "daemon.insecure",
			EnvVars:     []string{"PLUGIN_INSECURE"},
			Usage:       "docker daemon allows insecure registries",
			Destination: &settings.Daemon.Insecure,
		},
		&cli.BoolFlag{
			Name:        "daemon.ipv6",
			EnvVars:     []string{"PLUGIN_IPV6"},
			Usage:       "docker daemon IPv6 networking",
			Destination: &settings.Daemon.IPv6,
		},
		&cli.BoolFlag{
			Name:        "daemon.experimental",
			EnvVars:     []string{"PLUGIN_EXPERIMENTAL"},
			Usage:       "docker daemon Experimental mode",
			Destination: &settings.Daemon.Experimental,
		},
		&cli.BoolFlag{
			Name:        "daemon.debug",
			EnvVars:     []string{"PLUGIN_DEBUG", "DOCKER_LAUNCH_DEBUG"},
			Usage:       "docker daemon executes in debug mode",
			Destination: &settings.Daemon.Debug,
		},
		&cli.BoolFlag{
			Name:        "daemon.off",
			EnvVars:     []string{"PLUGIN_DAEMON_OFF"},
			Usage:       "don't start the docker daemon",
			Destination: &settings.Daemon.Disabled,
		},
		&cli.StringFlag{
			Name:        "daemon.buildkit-config",
			EnvVars:     []string{"PLUGIN_BUILDKIT_CONFIG"},
			Usage:       "docker buildkit json config content",
			Destination: &settings.Daemon.BuildkitConfig,
		},
		&cli.StringFlag{
			Name:        "dockerfile",
			EnvVars:     []string{"PLUGIN_DOCKERFILE"},
			Usage:       "build dockerfile",
			Value:       "Dockerfile",
			Destination: &settings.Build.Dockerfile,
		},
		&cli.StringFlag{
			Name:        "context",
			EnvVars:     []string{"PLUGIN_CONTEXT"},
			Usage:       "build context",
			Value:       ".",
			Destination: &settings.Build.Context,
		},
		&cli.StringSliceFlag{
			Name:        "tags",
			EnvVars:     []string{"PLUGIN_TAG", "PLUGIN_TAGS"},
			Usage:       "build tags",
			Value:       cli.NewStringSlice([]string{"latest"}...),
			FilePath:    ".tags",
			Destination: &settings.Build.Tags,
		},
		&cli.BoolFlag{
			Name:        "tags.auto",
			EnvVars:     []string{"PLUGIN_DEFAULT_TAGS", "PLUGIN_AUTO_TAG"},
			Usage:       "default build tags",
			Destination: &settings.Build.TagsAuto,
		},
		&cli.StringFlag{
			Name:        "tags.suffix",
			EnvVars:     []string{"PLUGIN_DEFAULT_SUFFIX", "PLUGIN_AUTO_TAG_SUFFIX"},
			Usage:       "default build tags with suffix",
			Destination: &settings.Build.TagsSuffix,
		},
		&cli.StringSliceFlag{
			Name:        "args",
			EnvVars:     []string{"PLUGIN_BUILD_ARGS"},
			Usage:       "build args",
			Destination: &settings.Build.Args,
		},
		&cli.StringSliceFlag{
			Name:        "args-from-env",
			EnvVars:     []string{"PLUGIN_BUILD_ARGS_FROM_ENV"},
			Usage:       "build args",
			Destination: &settings.Build.ArgsEnv,
		},
		&cli.BoolFlag{
			Name:        "quiet",
			EnvVars:     []string{"PLUGIN_QUIET"},
			Usage:       "quiet docker build",
			Destination: &settings.Build.Quiet,
		},
		&cli.StringFlag{
			Name:        "target",
			EnvVars:     []string{"PLUGIN_TARGET"},
			Usage:       "build target",
			Destination: &settings.Build.Target,
		},
		&cli.StringSliceFlag{
			Name:        "cache-from",
			EnvVars:     []string{"PLUGIN_CACHE_FROM"},
			Usage:       "images to consider as cache sources",
			Destination: &settings.Build.CacheFrom,
		},
		&cli.BoolFlag{
			Name:        "pull-image",
			EnvVars:     []string{"PLUGIN_PULL_IMAGE"},
			Usage:       "force pull base image at build time",
			Value:       true,
			Destination: &settings.Build.Pull,
		},
		&cli.BoolFlag{
			Name:        "compress",
			EnvVars:     []string{"PLUGIN_COMPRESS"},
			Usage:       "compress the build context using gzip",
			Destination: &settings.Build.Compress,
		},
		&cli.StringFlag{
			Name:        "repo",
			EnvVars:     []string{"PLUGIN_REPO"},
			Usage:       "docker repository",
			Destination: &settings.Build.Repo,
		},
		&cli.StringFlag{
			Name:        "docker.registry",
			EnvVars:     []string{"PLUGIN_REGISTRY", "DOCKER_REGISTRY"},
			Usage:       "docker registry",
			Value:       "https://index.docker.io/v1/",
			Destination: &settings.Login.Registry,
		},
		&cli.StringFlag{
			Name:        "docker.username",
			EnvVars:     []string{"PLUGIN_USERNAME", "DOCKER_USERNAME"},
			Usage:       "docker username",
			Destination: &settings.Login.Username,
		},
		&cli.StringFlag{
			Name:        "docker.password",
			EnvVars:     []string{"PLUGIN_PASSWORD", "DOCKER_PASSWORD"},
			Usage:       "docker password",
			Destination: &settings.Login.Password,
		},
		&cli.StringFlag{
			Name:        "docker.email",
			EnvVars:     []string{"PLUGIN_EMAIL", "DOCKER_EMAIL"},
			Usage:       "docker email",
			Destination: &settings.Login.Email,
		},
		&cli.StringFlag{
			Name:        "docker.config",
			EnvVars:     []string{"PLUGIN_CONFIG", "DOCKER_PLUGIN_CONFIG"},
			Usage:       "docker json dockerconfig content",
			Destination: &settings.Login.Config,
		},
		&cli.BoolFlag{
			Name:        "docker.purge",
			EnvVars:     []string{"PLUGIN_PURGE"},
			Usage:       "docker should cleanup images",
			Value:       true,
			Destination: &settings.Cleanup,
		},
		&cli.BoolFlag{
			Name:        "no-cache",
			EnvVars:     []string{"PLUGIN_NO_CACHE"},
			Usage:       "do not use cached intermediate containers",
			Destination: &settings.Build.NoCache,
		},
		&cli.StringSliceFlag{
			Name:        "add-host",
			EnvVars:     []string{"PLUGIN_ADD_HOST"},
			Usage:       "additional host:IP mapping",
			Destination: &settings.Build.AddHost,
		},
		&cli.StringSliceFlag{
			Name:        "platforms",
			EnvVars:     []string{"PLUGIN_PLATFORMS"},
			Usage:       "arget platform for build",
			Destination: &settings.Build.Platforms,
		},
	}
}
