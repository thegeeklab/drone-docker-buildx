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
			Usage:       "dry run disables docker push",
			EnvVars:     []string{"PLUGIN_DRY_RUN"},
			Destination: &settings.Dryrun,
		},
		&cli.StringFlag{
			Name:        "remote.url",
			Usage:       "git remote url",
			EnvVars:     []string{"DRONE_REMOTE_URL"},
			Destination: &settings.Build.Remote,
		},
		&cli.StringFlag{
			Name:        "daemon.mirror",
			Usage:       "docker daemon registry mirror",
			EnvVars:     []string{"PLUGIN_MIRROR", "DOCKER_PLUGIN_MIRROR"},
			Destination: &settings.Daemon.Mirror,
		},
		&cli.StringFlag{
			Name:        "daemon.storage-driver",
			Usage:       "docker daemon storage driver",
			EnvVars:     []string{"PLUGIN_STORAGE_DRIVER"},
			Destination: &settings.Daemon.StorageDriver,
		},
		&cli.StringFlag{
			Name:        "daemon.storage-path",
			Usage:       "docker daemon storage path",
			Value:       "/var/lib/docker",
			EnvVars:     []string{"PLUGIN_STORAGE_PATH"},
			Destination: &settings.Daemon.StoragePath,
		},
		&cli.StringFlag{
			Name:        "daemon.bip",
			Usage:       "docker daemon bride ip address",
			EnvVars:     []string{"PLUGIN_BIP"},
			Destination: &settings.Daemon.Bip,
		},
		&cli.StringFlag{
			Name:        "daemon.mtu",
			Usage:       "docker daemon custom mtu setting",
			EnvVars:     []string{"PLUGIN_MTU"},
			Destination: &settings.Daemon.MTU,
		},
		&cli.StringSliceFlag{
			Name:        "daemon.dns",
			Usage:       "docker daemon dns server",
			EnvVars:     []string{"PLUGIN_CUSTOM_DNS"},
			Destination: &settings.Daemon.DNS,
		},
		&cli.StringSliceFlag{
			Name:        "daemon.dns-search",
			Usage:       "docker daemon dns search domains",
			EnvVars:     []string{"PLUGIN_CUSTOM_DNS_SEARCH"},
			Destination: &settings.Daemon.DNSSearch,
		},
		&cli.BoolFlag{
			Name:        "daemon.insecure",
			Usage:       "docker daemon allows insecure registries",
			EnvVars:     []string{"PLUGIN_INSECURE"},
			Destination: &settings.Daemon.Insecure,
		},
		&cli.BoolFlag{
			Name:        "daemon.ipv6",
			Usage:       "docker daemon IPv6 networking",
			EnvVars:     []string{"PLUGIN_IPV6"},
			Destination: &settings.Daemon.IPv6,
		},
		&cli.BoolFlag{
			Name:        "daemon.experimental",
			Usage:       "docker daemon Experimental mode",
			EnvVars:     []string{"PLUGIN_EXPERIMENTAL"},
			Destination: &settings.Daemon.Experimental,
		},
		&cli.BoolFlag{
			Name:        "daemon.debug",
			Usage:       "docker daemon executes in debug mode",
			EnvVars:     []string{"PLUGIN_DEBUG", "DOCKER_LAUNCH_DEBUG"},
			Destination: &settings.Daemon.Debug,
		},
		&cli.BoolFlag{
			Name:        "daemon.off",
			Usage:       "don't start the docker daemon",
			EnvVars:     []string{"PLUGIN_DAEMON_OFF"},
			Destination: &settings.Daemon.Disabled,
		},
		&cli.StringFlag{
			Name:        "daemon.buildkit-config",
			Usage:       "location of buildkit config file",
			EnvVars:     []string{"PLUGIN_BUILDKIT_CONFIG"},
			Destination: &settings.Daemon.BuildkitConfig,
		},
		&cli.StringFlag{
			Name:        "dockerfile",
			Usage:       "build dockerfile",
			Value:       "Dockerfile",
			EnvVars:     []string{"PLUGIN_DOCKERFILE"},
			Destination: &settings.Build.Dockerfile,
		},
		&cli.StringFlag{
			Name:        "context",
			Usage:       "build context",
			Value:       ".",
			EnvVars:     []string{"PLUGIN_CONTEXT"},
			Destination: &settings.Build.Context,
		},
		&cli.StringSliceFlag{
			Name:        "tags",
			Usage:       "build tags",
			Value:       cli.NewStringSlice([]string{"latest"}...),
			EnvVars:     []string{"PLUGIN_TAG", "PLUGIN_TAGS"},
			FilePath:    ".tags",
			Destination: &settings.Build.Tags,
		},
		&cli.BoolFlag{
			Name:        "tags.auto",
			Usage:       "default build tags",
			EnvVars:     []string{"PLUGIN_DEFAULT_TAGS", "PLUGIN_AUTO_TAG"},
			Destination: &settings.Build.TagsAuto,
		},
		&cli.StringFlag{
			Name:        "tags.suffix",
			Usage:       "default build tags with suffix",
			EnvVars:     []string{"PLUGIN_DEFAULT_SUFFIX", "PLUGIN_AUTO_TAG_SUFFIX"},
			Destination: &settings.Build.TagsSuffix,
		},
		&cli.StringSliceFlag{
			Name:        "args",
			Usage:       "build args",
			EnvVars:     []string{"PLUGIN_BUILD_ARGS"},
			Destination: &settings.Build.Args,
		},
		&cli.StringSliceFlag{
			Name:        "args-from-env",
			Usage:       "build args",
			EnvVars:     []string{"PLUGIN_BUILD_ARGS_FROM_ENV"},
			Destination: &settings.Build.ArgsEnv,
		},
		&cli.BoolFlag{
			Name:        "quiet",
			Usage:       "quiet docker build",
			EnvVars:     []string{"PLUGIN_QUIET"},
			Destination: &settings.Build.Quiet,
		},
		&cli.StringFlag{
			Name:        "target",
			Usage:       "build target",
			EnvVars:     []string{"PLUGIN_TARGET"},
			Destination: &settings.Build.Target,
		},
		&cli.StringSliceFlag{
			Name:        "cache-from",
			Usage:       "images to consider as cache sources",
			EnvVars:     []string{"PLUGIN_CACHE_FROM"},
			Destination: &settings.Build.CacheFrom,
		},
		&cli.BoolFlag{
			Name:        "squash",
			Usage:       "squash the layers at build time",
			EnvVars:     []string{"PLUGIN_SQUASH"},
			Destination: &settings.Build.Squash,
		},
		&cli.BoolFlag{
			Name:        "pull-image",
			Usage:       "force pull base image at build time",
			EnvVars:     []string{"PLUGIN_PULL_IMAGE"},
			Value:       true,
			Destination: &settings.Build.Pull,
		},
		&cli.BoolFlag{
			Name:        "compress",
			Usage:       "compress the build context using gzip",
			EnvVars:     []string{"PLUGIN_COMPRESS"},
			Destination: &settings.Build.Compress,
		},
		&cli.StringFlag{
			Name:        "repo",
			Usage:       "docker repository",
			EnvVars:     []string{"PLUGIN_REPO"},
			Destination: &settings.Build.Repo,
		},
		&cli.StringFlag{
			Name:        "docker.registry",
			Usage:       "docker registry",
			Value:       "https://index.docker.io/v1/",
			EnvVars:     []string{"PLUGIN_REGISTRY", "DOCKER_REGISTRY"},
			Destination: &settings.Login.Registry,
		},
		&cli.StringFlag{
			Name:        "docker.username",
			Usage:       "docker username",
			EnvVars:     []string{"PLUGIN_USERNAME", "DOCKER_USERNAME"},
			Destination: &settings.Login.Username,
		},
		&cli.StringFlag{
			Name:        "docker.password",
			Usage:       "docker password",
			EnvVars:     []string{"PLUGIN_PASSWORD", "DOCKER_PASSWORD"},
			Destination: &settings.Login.Password,
		},
		&cli.StringFlag{
			Name:        "docker.email",
			Usage:       "docker email",
			EnvVars:     []string{"PLUGIN_EMAIL", "DOCKER_EMAIL"},
			Destination: &settings.Login.Email,
		},
		&cli.StringFlag{
			Name:        "docker.config",
			Usage:       "docker json dockerconfig content",
			EnvVars:     []string{"PLUGIN_CONFIG", "DOCKER_PLUGIN_CONFIG"},
			Destination: &settings.Login.Config,
		},
		&cli.BoolFlag{
			Name:        "docker.purge",
			Usage:       "docker should cleanup images",
			EnvVars:     []string{"PLUGIN_PURGE"},
			Value:       true,
			Destination: &settings.Cleanup,
		},
		&cli.BoolFlag{
			Name:        "no-cache",
			Usage:       "do not use cached intermediate containers",
			EnvVars:     []string{"PLUGIN_NO_CACHE"},
			Destination: &settings.Build.NoCache,
		},
		&cli.StringSliceFlag{
			Name:        "add-host",
			Usage:       "additional host:IP mapping",
			EnvVars:     []string{"PLUGIN_ADD_HOST"},
			Destination: &settings.Build.AddHost,
		},
		&cli.StringSliceFlag{
			Name:        "platforms",
			Usage:       "arget platform for build",
			EnvVars:     []string{"PLUGIN_PLATFORMS"},
			Destination: &settings.Build.Platforms,
		},
	}
}
